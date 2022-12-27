#[derive(Clone)]
pub struct Rule {
    before: Vec<char>,
    after: Vec<char>,
    is_terminate: bool,
}

impl Rule {
    pub fn new(before: &str, after: &str, is_terminate: bool) -> Self {
        Rule {
            before: before.chars().collect(),
            after: after.chars().collect(),
            is_terminate,
        }
    }
}

pub struct Markov {
    rules: Vec<Rule>,
    text: Vec<char>,
}

impl Markov {
    pub fn new(code: &str) -> Result<Self, &str> {
        match Markov::parse(code) {
            Ok(rules) => Ok(Markov {
                rules: rules,
                text: Vec::<char>::new(),
            }),
            Err(msg) => Err(msg),
        }
    }

    fn parse(code: &str) -> Result<Vec<Rule>, &str> {
        let code_chars: Vec<char> = code.chars().collect();
        let mut rules = vec![];
        let mut i = 0 as usize;
        while i < code_chars.len() {
            let mut before = vec![];
            let mut after = vec![];
            let is_terminate;

            while i < code_chars.len() && code_chars[i] != ':' {
                before.push(code_chars[i]);
                i += 1;
            }

            if i < code_chars.len() && code_chars[i] == ':' {
                if i + 1 < code_chars.len() && code_chars[i + 1] == ':' {
                    is_terminate = true;
                    i += 2;
                } else {
                    is_terminate = false;
                    i += 1;
                }
            } else {
                return Err("Syntax Error");
            }

            while i < code_chars.len() && code_chars[i] != '\n' {
                after.push(code_chars[i]);
                i += 1;
            }

            rules.push(Rule {
                before,
                after,
                is_terminate,
            });

            i += 1;
        }
        return Ok(rules);
    }

    pub fn set_text(&mut self, text: &str) {
        self.text = text.chars().collect();
    }

    pub fn run(&mut self) -> Vec<char> {
        loop {
            let (next_text, is_terminate) = self.step();
            if is_terminate {
                return next_text;
            } else {
                self.text = next_text;
            }
        }
    }

    pub fn step(&mut self) -> (Vec<char>, bool) {
        for rule in self.rules.clone() {
            println!("{:?} {:?} {:?}", rule.before, rule.after, rule.is_terminate);
            if self.text.len() < rule.before.len() {
                continue;
            }

            for rule_start in (0 as usize)..((self.text.len() - rule.before.len() + 1) as usize) {
                let mut pattern_match = true;
                for i in (0 as usize)..(rule.before.len() as usize) {
                    if self.text[rule_start + i] != rule.before[i] {
                        pattern_match = false;
                        break;
                    }
                }

                if pattern_match {
                    let mut next_text: Vec<char> = vec![];
                    for i in 0..rule_start {
                        next_text.push(self.text[i]);
                    }
                    for c in rule.after {
                        next_text.push(c);
                    }
                    for i in rule_start + rule.before.len()..self.text.len() {
                        next_text.push(self.text[i]);
                    }
                    return (next_text, rule.is_terminate);
                }
            }
        }

        return (self.text.clone(), true);
    }
}

#[test]
fn markov_man_woman() {
    if let Ok(mut markov) = Markov::new("woman:W\nman:M\nMW:\nWM:\n") {
        markov.set_text("manmanwomanwomanmanwomanwomanmanwomanmanmanwoman");
        println!("{:?}", markov.run());
    }
}

#[test]
fn markov_divide_2() {
    if let Ok(mut markov) = Markov::new("s0:0s\ns1:1s\n0s::\n:s\n") {
        markov.set_text("1001010101010111010");
        println!("{:?}", markov.run());
    }
}
