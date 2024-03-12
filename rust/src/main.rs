use clap::{Parser, Subcommand};
#[derive(Parser)]
#[command(author,version,about,long_about=None)]
struct Cli {
    #[command(subcommand)]
    command: Option<Commands>,
}

#[derive(Subcommand)]
enum Commands {
    Genpwd {
        #[arg(short, long)]
        file: String,
    },
    Conn {
        #[arg(short, long)]
        server: String,
        #[arg(short, long)]
        port: String,
        #[arg(short, long)]
        user: String,
    },
}
fn main() {
    let cli = Cli::parse();

    match cli.command {
        Some(Commands::Genpwd { file }) => println!("{:?}", file),
        Some(Commands::Conn { server, port, user }) => {
            println!("server:{}", server);
            println!("port:{}", port);
            println!("user:{}", user);
        },
        None => {}
    }
}
