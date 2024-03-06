Last login: Sat Feb 24 17:01:22 on ttys000
You have new mail.
❯ cd Personal
cd: no such file or directory: Personal
❯ cd Desktop/Personal
❯ git clone https://github.com/IdealisticINTJ/coding-challenges.git

Cloning into 'coding-challenges'...
remote: Enumerating objects: 9, done.
remote: Counting objects: 100% (9/9), done.
remote: Compressing objects: 100% (6/6), done.
remote: Total 9 (delta 2), reused 0 (delta 0), pack-reused 0
Receiving objects: 100% (9/9), done.
Resolving deltas: 100% (2/2), done.
❯ cd coding-challenges

❯ mkdir challenge1

❯ cat << EOF > challenge1/main.go

heredoc> >....                                                                  
        flag.BoolVar(&countBytes, "c", false, "Count bytes")
        flag.BoolVar(&countLines, "l", false, "Count lines")
        flag.BoolVar(&countWords, "w", false, "Count words")
        flag.BoolVar(&countChars, "m", false, "Count characters")

        flag.Parse()

        if !countBytes && !countLines && !countWords && !countChars {
                countBytes = true
                countLines = true
                countWords = true
        }

        filenames := flag.Args()

        if len(filenames) == 0 {
                challenge1.CountFromStdin(countBytes, countLines, countWords, countChars)
        } else {
                challenge1.CountFromFiles(filenames, countBytes, countLines, countWords, countChars)
        }
}   
heredoc> EOF

❯ git add challenge1/main.go

❯ git commit -m "feat(challenge1): Implement command-line flag parsing and file handling in main"

[main 3855cd3] feat(challenge1): Implement command-line flag parsing and file handling in main
 1 file changed, 34 insertions(+)
 create mode 100644 challenge1/main.go
❯ cat << EOF > challenge1/stats.go

heredoc> >....                                                                  
1}},
                {"Multibyte chars", "s⌘ f", stats{bytes: 6, words: 2, lines: 0, chars: 4}},
                {"Trailing newline", "this is a sentence\n\nacross multiple\nlines\n", stats{bytes: 42, words: 7, lines: 4, chars: 42}},
                {"No trailing newline", "this is a sentence\n\nacross multiple\nlines", stats{bytes: 41, words: 7, lines: 3, chars: 41}},
        }

        for _, test := range cases {
                t.Run(test.Description, func(t *testing.T) {
                        bufferedString := bufio.NewReader(strings.NewReader(test.Input))
                        got := calculateStats(bufferedString, true, true, true, true)

                        if got != test.Want {
                                t.Errorf("got %v, want %v", got, test.Want)
                        }
                })
        }
}
EOF
❯ git add challenge1/stats.go

❯ git commit -m "feat(challenge1): Add functions for counting bytes, lines, words, and characters"

[main 719fe49] feat(challenge1): Add functions for counting bytes, lines, words, and characters
 1 file changed, 33 insertions(+)
 create mode 100644 challenge1/stats.go
❯ cat << EOF > challenge1/stats_test.go

heredoc> >....                                                                  
                InputFilename string
                Want          string
        }{
                {"Empty", stats{bytes: 0, words: 0, lines: 0, chars: 0}, "", "0\t0\t0\t"},
                {"Default", stats{bytes: 11, words: 2, lines: 1, chars: 0}, "filename", "1\t2\t11\tfilename"},
                {"Chars", stats{bytes: 0, words: 0, lines: 0, chars: 100}, "filename", "100\tfilename"},
        }

        for _, test := range cases {
                t.Run(test.Description, func(t *testing.T) {
                        got := formatStats(true, true, true, true, test.InputStats, test.InputFilename)

                        if got != test.Want {
                                t.Errorf("got %v, want %v", got, test.Want)
                        }
                })
        }
}
EOF
❯ 
git commit -m "test(challenge1): Enhance test coverage for stats"
On branch main
Your branch is ahead of 'origin/main' by 2 commits.
  (use "git push" to publish your local commits)

Untracked files:
  (use "git add <file>..." to include in what will be committed)
	challenge1/stats_test.go

nothing added to commit but untracked files present (use "git add" to track)
❯ git push origin main

Username for 'https://github.com': IdealisticINTJ
Password for 'https://IdealisticINTJ@github.com': 
remote: Support for password authentication was removed on August 13, 2021.
remote: Please see https://docs.github.com/get-started/getting-started-with-git/about-remote-repositories#cloning-with-https-urls for information on currently recommended modes of authentication.
fatal: Authentication failed for 'https://github.com/IdealisticINTJ/coding-challenges.git/'
❯ git push origin main

Username for 'https://github.com': IdealisticINTJ
Password for 'https://IdealisticINTJ@github.com': 
remote: Permission to IdealisticINTJ/coding-challenges.git denied to IdealisticINTJ.
fatal: unable to access 'https://github.com/IdealisticINTJ/coding-challenges.git/': The requested URL returned error: 403
❯ git push origin main

Username for 'https://github.com': IdealisticINTJ
Password for 'https://IdealisticINTJ@github.com': 
remote: Invalid username or password.
fatal: Authentication failed for 'https://github.com/IdealisticINTJ/coding-challenges.git/'
❯ git push origin main

Username for 'https://github.com': IdealisticINTJ
Password for 'https://IdealisticINTJ@github.com': 
Enumerating objects: 9, done.
Counting objects: 100% (9/9), done.
Delta compression using up to 8 threads
Compressing objects: 100% (7/7), done.
Writing objects: 100% (8/8), 1.43 KiB | 1.43 MiB/s, done.
Total 8 (delta 0), reused 0 (delta 0), pack-reused 0
To https://github.com/IdealisticINTJ/coding-challenges.git
   d50fd0b..719fe49  main -> main
❯ cat << EOF > challenge1/stats_test.go

heredoc> >....                                                                                                                   

func TestFormatStats(t *testing.T) {
        cases := []struct {
                Description   string
                InputStats    stats
                InputFilename string
                Want          string
        }{
                {"Empty", stats{bytes: 0, words: 0, lines: 0, chars: 0}, "", "0\t0\t0\t"},
                {"Default", stats{bytes: 11, words: 2, lines: 1, chars: 0}, "filename", "1\t2\t11\tfilename"},
                {"Chars", stats{bytes: 0, words: 0, lines: 0, chars: 100}, "filename", "100\tfilename"},
        }

        for _, test := range cases {
                t.Run(test.Description, func(t *testing.T) {
                        got := formatStats(true, true, true, true, test.InputStats, test.InputFilename)

                        if got != test.Want {
                                t.Errorf("got %v, want %v", got, test.Want)
                        }
                })
        }
}
heredoc> EOF
❯ git add challenge1/stats_test.go

❯ git commit -m "test(challenge1): Enhance test coverage for stats"

[main de24f6e] test(challenge1): Enhance test coverage for stats
 1 file changed, 27 insertions(+)
 create mode 100644 challenge1/stats_test.go
❯ git push origin main

Enumerating objects: 6, done.
Counting objects: 100% (6/6), done.
Delta compression using up to 8 threads
Compressing objects: 100% (4/4), done.
Writing objects: 100% (4/4), 743 bytes | 743.00 KiB/s, done.
Total 4 (delta 0), reused 0 (delta 0), pack-reused 0
To https://github.com/IdealisticINTJ/coding-challenges.git
   719fe49..de24f6e  main -> main
❯ mkdir challenge2

❯ cd challenge2

❯ touch main.go

❯ nano main.go

  UW PICO 5.09                                              File: main.go                                              Modified  

        }
                        
        if firstToken.Type != TokenObjectStart || lastToken.Type != TokenObjectEnd {
                fmt.Println("Invalid JSON")
                os.Exit(1)
        }
                        
        fmt.Println("Valid JSON")
}               
                
                
                        
        
                
                
                        
                        
        
                

^G Get Help          ^O WriteOut          ^R Read File         ^Y Prev Pg           ^K Cut Text          ^C Cur Pos           
^X Exit              ^J Justify           ^W Where is          ^V Next Pg           ^U UnCut Text        ^T To Spell          
