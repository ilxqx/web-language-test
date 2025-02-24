use std::sync::Arc;

use axum::{Router, extract::State, response::Json, routing::get};
use sea_orm::{ColumnTrait, DatabaseConnection, EntityTrait, QueryFilter};
use tokio::net::TcpListener;
use tracing::info;

use crate::{
    AppResult, SysUsers,
    entity::{Column, Model as SysUsersModel},
};

pub struct AppState {
    db_connection: DatabaseConnection,
}

impl AppState {
    pub fn new(db_connection: DatabaseConnection) -> Self {
        Self { db_connection }
    }
}

pub struct App {
    db_connection: DatabaseConnection,
}

impl App {
    pub fn new(db_connection: DatabaseConnection) -> Self {
        Self { db_connection }
    }

    pub async fn run(&self) -> AppResult<()> {
        let app = Router::new()
            .route("/", get(index))
            .with_state(Arc::new(AppState::new(self.db_connection.clone())));
        let listener = TcpListener::bind("0.0.0.0:3000").await?;

        info!("Server is running on http://0.0.0.0:3000");
        axum::serve(listener, app).await?;
        Ok(())
    }
}

async fn index(State(state): State<Arc<AppState>>) -> Json<Vec<SysUsersModel>> {
    let random_id1 = fastrand::i32(0..1000);
    let random_id2 = fastrand::i32(8000..9000);
    info!("Generated random id1: {}, id2: {}", random_id1, random_id2);

    let condition = Column::UserId.between(random_id1, random_id2);
    match SysUsers::find()
        .filter(condition)
        .all(&state.db_connection)
        .await
    {
        Ok(users) => {
            info!("Found {} users", users.len());
            Json(users)
        }
        Err(err) => {
            info!("Error: {}", err);
            Json(Vec::new())
        }
    }
}
