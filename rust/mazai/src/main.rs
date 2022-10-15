use tide::{Body, Request, Response};
use tide::prelude::*;
use tera::Tera;
use tide_tera::prelude::*;
use serde_json::{Number, Value};

#[derive(Debug, Deserialize, Serialize)]
struct Animal {
    name: String,
    legs: u16,
}

#[derive(Debug, Deserialize, Serialize)]
struct TripDay {
    seats: u16,
    guests: u16
}

#[derive(Debug, Deserialize, Serialize)]
struct Trip {
    to: String,
    from: String,
    desc: String,
    driver: String,
    time: String,
    delay: u16,
    days: [TripDay;7]
}

struct State {
    tera: Tera,
    trips: Vec<Trip>
}


#[async_std::main]
async fn main() -> tide::Result<()> {
    let mut ter = Tera::new("templates/**/*")?;
    ter.autoescape_on(vec!["html"]);
    //let mut state = State{tera: ter, trips: vec![]};

    let mut app = tide::with_state(ter); // tide::new()

    //app.at("/orders/shoes").post(order_shoes);
    //app.at("/").get(info);
    // app.at("/:name").get(|req: tide::Request<Tera>| async move {
    //     let tera = req.state();
    //     let name: String = req.param("name")?.to_owned();
    //     tera.render_response("index.html", &context! { "name" => name })
    // });
    app.at("/").get(|req: tide::Request<Tera>| async move {
        let tera = req.state();
        tera.render_response("index.html", &context! {})
    });

    app.at("/trips").get(|req: tide::Request<Tera>| async move {

        let data = std::fs::read_to_string("trips.json").expect("Unable to read file");
        let json: serde_json::Value = serde_json::from_str(&data).expect("JSON does not have correct format.");

        /*let trips = [
            Trip {
                to: "1".into(), from: "2".into(), desc: "hoho".into(), driver: "me".into(), 
                time: "10:30".into(), delay: 10, days: [
                    TripDay{seats:5,guests:3}, TripDay{seats:5,guests:3}, TripDay{seats:5,guests:3},
                    TripDay{seats:5,guests:3}, TripDay{seats:0,guests:0}, TripDay{seats:0,guests:0},
                    TripDay{seats:0,guests:0}
                ]
            }
        ];*/

        //let mut res = tide::Response::new(200).set_body(Body::from_json(trip);
        //res.set_body(Body::from_json(trip)?);
        //.body_json(&cat).unwrap()

        //let tera = req.state();
        //tera.render_response("index.html", &context! {})

        let mut res = tide::Response::new(200);
        res.set_body(Body::from_json(&json)?);
        Ok(res)
    });

    app.at("/trip/:id/edit").get( |req: tide::Request<Tera> | async move {
        let tera = req.state();
        let strId = req.param("id");
        let id: i32 = strId.parse().expect("Can't parse int");
        tera.render_response("trip.html", &context! {"id" => id})
    });

    app.at("/trip/:id").get( |req: tide::Request<Tera> | async move {
        let trip =
            Trip {
                to: "1".into(), from: "2".into(), desc: "hoho".into(), driver: "me".into(), 
                time: "10:30".into(), delay: 10, days: [
                    TripDay{seats:5,guests:3}, TripDay{seats:5,guests:3}, TripDay{seats:5,guests:3},
                    TripDay{seats:5,guests:3}, TripDay{seats:0,guests:0}, TripDay{seats:0,guests:0},
                    TripDay{seats:0,guests:0}
                ]
            };

        let mut res = tide::Response::new(200);
        res.set_body(Body::from_json(&trip)?);
        Ok(res)
    });

    // Serve static files
    app.at("/static").serve_dir("static/").expect("Invalid static file directory");

    app.listen("127.0.0.1:8080").await?;

    println!("Server is closed.");
    Ok(())
}

async fn order_shoes(mut req: Request<()>) -> tide::Result {
    let Animal { name, legs } = req.body_json().await?;
    Ok(format!("Hello, {}! I've put in an order for {} shoes", name, legs).into())
}

async fn info(_req: Request<()>) -> tide::Result {
    Ok(format!("Hello\n").into())
}

/*
async fn GetHomePage(_req: Request<()>) -> tide::Result {
	http.ServeFile(w, r, "templates/index.html")
}
*/