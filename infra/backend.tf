terraform {
  backend "gcs" {
    bucket = "finmng-service-api-tfstate"
    prefix = "env/dev"
  }
}
