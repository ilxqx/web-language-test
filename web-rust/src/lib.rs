use app::App;
use sea_orm::{ConnectOptions, Database, DatabaseConnection};
use tracing::info;
use tracing::log::LevelFilter;
use tracing::{Level, subscriber::set_global_default};
use tracing_subscriber::fmt::Subscriber;

pub mod app;
pub mod entity;

pub use entity::Entity as SysUsers;
pub use entity::Model as SysUsersModel;

pub type AppResult<T> = Result<T, anyhow::Error>;

/// Initialize the tracing subscriber
pub fn init_tracing() {
    let subscriber = Subscriber::builder()
        .with_max_level(Level::INFO)
        .with_line_number(true)
        .with_file(true)
        .with_target(false)
        .finish();

    set_global_default(subscriber).expect("Failed to set default subscriber");
}

/// Creates a connection to the database.
pub async fn create_database_connection() -> AppResult<DatabaseConnection> {
    let mut options = ConnectOptions::new("postgres://postgres:12345678@127.0.0.1:5432/postgres");
    options
        .max_connections(200)
        .sqlx_logging(false)
        .sqlx_logging_level(LevelFilter::Info);

    Ok(Database::connect(options).await?)
}

pub async fn start_server() -> AppResult<()> {
    init_tracing();
    let db_connection = create_database_connection().await?;
    info!("Connected to database successfully");
    let app = App::new(db_connection);
    app.run().await?;
    Ok(())
}
