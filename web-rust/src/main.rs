use web_rust::{AppResult, start_server};

#[tokio::main]
async fn main() -> AppResult<()> {
    start_server().await
}
