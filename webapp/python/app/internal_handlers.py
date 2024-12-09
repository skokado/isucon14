from http import HTTPStatus

from fastapi import APIRouter
from sqlalchemy import text

from app.models import Chair, Ride
from app.sql import engine

router = APIRouter(prefix="/api/internal")


# このAPIをインスタンス内から一定間隔で叩かせることで、椅子とライドをマッチングさせる
@router.get("/matching", status_code=HTTPStatus.NO_CONTENT)
def internal_get_matching() -> None:
    with engine.begin() as conn:
        waiting_ride_row = conn.execute(
            text(
                "SELECT * FROM rides WHERE chair_id IS NULL ORDER BY created_at LIMIT 1"
            )
        ).fetchone()
    if waiting_ride_row is None:
        return
    ride = Ride.model_validate(waiting_ride_row)

    ignore_chair_ids = []
    for _ in range(10):
        with engine.begin() as conn:
            active_chair_rows = conn.execute(
                text("""
                    SELECT
                        id,
                        current_latitude,
                        current_longitude
                    FROM chairs
                    WHERE is_active = TRUE
                      AND current_latitude IS NOT NULL
                      AND id NOT IN (:ignore_chair_ids)
                """),
                {"ignore_chair_ids": ",".join(ignore_chair_ids) if ignore_chair_ids else "''"},
            ).fetchall()
        if active_chair_rows is None:
            return

        distances = [
            (
                row.id,
                abs(ride.pickup_latitude - row[1]),
                abs(ride.pickup_longitude - row[2]),
            )
            for row in active_chair_rows
            if row[1] is not None and row[2] is not None
        ]
        nearest_chair_id = min(distances, key=lambda x: x[1] + x[2])[0]
        
        with engine.begin() as conn:
            if can_assign := bool(
                conn.execute(
                    text("""
                        SELECT
                            COUNT(*) = 0
                        FROM
                            (
                            SELECT
                                COUNT(chair_sent_at) = 6 AS completed
                            FROM
                                ride_statuses
                            WHERE
                                ride_id IN (
                                SELECT
                                    id
                                FROM
                                    rides
                                WHERE
                                    chair_id = :chair_id)
                            GROUP BY
                                ride_id) is_completed
                        WHERE
                            completed = FALSE 
                    """),
                    {"chair_id": nearest_chair_id},
                ).scalar()
            ):
                break
            else:
                ignore_chair_ids.append(nearest_chair_id)
                continue
    else:
        return

    with engine.begin() as conn:
        conn.execute(
            text("UPDATE rides SET chair_id = :chair_id WHERE id = :id"),
            {"chair_id": nearest_chair_id, "id": ride.id},
        )
