use reqwest::Client;

async fn fetch_data() -> Result<String, Box<dyn std::error::Error>> {
    let client = Client::new();
    let response = client.get("https://www.fitcalcs.life").send().await?;

    if response.status().is_success() {
        let data = response.text().await?;
        Ok(data)
    } else {
        Err(format!("* * * Error fetching data: {}", response.status()).into())
    }
}