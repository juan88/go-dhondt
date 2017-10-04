package dhondt

import ("strings"
    "bufio"
    "strconv"
    "regexp"
    "os"
)

type ElectionResults struct {
	Parties uint
	Votes map[string]uint64
	Seats uint
}

func (res *ElectionResults) LoadResults(fname string, delim string) (err error) {

    res.Parties = 0
    file, err := os.Open(fname)

    if err == nil {

        // Builds two regular expressions to try and match them against
        // lines read using Go's bufio.Scanner.

        // TODO: find a better pattern to describe party names, since
        //       [[:print:]]+ may match against a completely blank string!

        limbs := []string { "[[:print:]]+",
            "[[:blank:]]*[[:digit:]]+[[:blank:]]*$" }

        // `delim' can be a string; that is, it can be comprised of more than
        // one ASCII printable character, including regexp metacharacters like `.'
        // which, if present, will be appropriately escaped by Go's `regexp.QuoteMeta'.
        
        // Comments.
        comment, _ := regexp.Compile("[[:blank:]]*#.*")

        // Valid data lines.
        line, err := regexp.Compile(strings.Join(limbs, regexp.QuoteMeta(delim)))

        if err == nil {

            res.Votes = make(map[string]uint64)
            res.Seats = 0

            scanner := bufio.NewScanner(file)

            for scanner.Scan() {

                byteLine := []byte(scanner.Text())

                if comment.Match(byteLine) {
                    continue // Ignores the current and reads the next line.
                } else if match := line.Find(byteLine); match != nil {

                    // Due to the fact that a valid data line can contain the delimiter
                    // as part of the party's name, the matched string should be splitted,
                    // ideally, into exactly two tokens at the last occurrence of
                    // `delim'. Instead of doing that, here we split the whole string
                    // at __every__ occurence of `delim' into a set of substrings
                    // and then re-join all but the last one of them, appropriately
                    // re-interposing the removed `delim' string.

                    fields := strings.Split(string(match), delim)
                    lastidx := len(fields) - 1

                    res.Votes[strings.Join(fields[:lastidx], delim)], _ =
                        strconv.ParseUint(fields[lastidx], 10, 64)
                    res.Parties++
                }                
            }
        }

        err = file.Close()
    }

    return
}
