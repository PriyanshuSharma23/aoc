use std::{cmp, fs, io::Read};

fn main() {
    let mut file = fs::File::open("input.txt").unwrap();
    let mut data = String::new();
    let _ = file.read_to_string(&mut data);


    let output = solve(&data);
    println!("{output}");
}

fn solve(data: &str) -> u32 {
    // let mut ref_map = HashMap::new();
    // ref_map.insert("red", 12);
    // ref_map.insert("green", 13);
    // ref_map.insert("blue", 14);

    return data
        .lines()
        .map(|line| {
            // let mut valid = true;
            let mut splits = line.split(":");
            splits.next();
            let games = splits.next().unwrap();
            let (mut red, mut green, mut blue) = (0, 0, 0);

            for game in games.split(";") {
                for set in game.split(",") {
                    println!("{set}");
                    let mut temp = set.trim().split(" ");
                    let number: u32 = temp.next().unwrap().parse().unwrap();
                    let color = temp.next().unwrap();

                    // match ref_map.get(&color).unwrap().cmp(&number) {
                    //     std::cmp::Ordering::Less => {
                    //         valid = false;
                    //         break;
                    //     }
                    //     _ => (),
                    // }
                    match color {
                        "red" => {
                            red = cmp::max(red, number);
                        }
                        "green" => {
                            green = cmp::max(green, number);
                        }
                        "blue" => {
                            blue = cmp::max(blue, number);
                        }
                        _ => (),
                    }
                }
                // if !valid { break }
            }

            // return if valid { idx as u32 + 1 } else { 0 };
            return red * green * blue;
        })
        .sum();
}
