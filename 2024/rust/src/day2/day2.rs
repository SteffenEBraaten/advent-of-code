use std::fs::File;
use std::io::Read;

pub fn part_1() -> i32 {
    let mut file = File::open("../input/day2.txt").expect("Could not open file");
    let mut contents = String::new();

    file.read_to_string(&mut contents)
        .expect("Could not convert file to string");

    let mut num_of_safe_reports = 0;

    for line in contents.lines() {
        let mut levels = Vec::new();

        for token in line.split_whitespace() {
            if let Ok(digit) = token.parse::<i32>() {
                levels.push(digit);
            }
        }

        let mut all_ascending = false;
        let mut all_descending = false;
        let mut within_threshold = true;

        for level in levels.windows(2) {
            if let [a, b] = level {
                let diff = (b - a).abs();
                if diff > 3 || diff < 1 {
                    within_threshold = false;
                }
                if b > a {
                    all_descending = true;
                }
                if b < a {
                    all_ascending = true;
                }
            }
        }

        if within_threshold && (all_ascending != all_descending) {
            num_of_safe_reports += 1;
        }
    }

    num_of_safe_reports
}
