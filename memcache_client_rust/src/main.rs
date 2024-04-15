use actix_web::{web, App, HttpServer, Responder};
use memcache::Client;
use std::sync::Arc;
use tokio::sync::Mutex;

async fn set_key(
    key: web::Path<String>,
    query_params: web::Query<std::collections::HashMap<String, String>>,
    client: web::Data<Arc<Mutex<Client>>>,
) -> impl Responder {
    if let Some(value) = query_params.get("value") {
        let value = value.clone(); // Clone the value to own it
        let mut client = client.lock().await;
        client.set(&key, &*value, 0).unwrap(); // Dereference the value before passing
        format!("Key '{}' set successfully", key)
    } else {
        "Value parameter not provided".to_string()
    }
}

async fn get_key(
    key: web::Path<String>,
    client: web::Data<Arc<Mutex<Client>>>,
) -> impl Responder {
    let key = key.into_inner();
    let mut client = client.lock().await;
    match client.get::<String>(&key) {
        Ok(Some(v)) => format!("Value for key '{}': {:?}", key, v),
        _ => format!("Key '{}' not found", key),
    }
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    let client = Client::connect("memcache://memcached:11211").unwrap();
    let client = Arc::new(Mutex::new(client));

    HttpServer::new(move || {
        App::new()
            .app_data(web::Data::new(client.clone()))
            .route("/set/{key}", web::post().to(set_key))
            .route("/get/{key}", web::get().to(get_key))
    })
    .bind("0.0.0.0:8080")?
    .run()
    .await
}