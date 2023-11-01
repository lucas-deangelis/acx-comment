# Uriage

Astral Codex Ten comment scrapper.

Get the articles:

```shell
curl 'https://www.astralcodexten.com/api/v1/archive?sort=new&search=&offset=0&limit=12' \
  -H 'authority: www.astralcodexten.com' \
  -H 'accept: */*' \
  -H 'accept-language: en-GB,en;q=0.9,fr-FR;q=0.8,fr;q=0.7,en-US;q=0.6,ja;q=0.5' \
  -H 'cookie: ajs_anonymous_id=%225a9f382a-701a-4ee7-97b4-d9eb45d47c4e%22; ajs_anonymous_id=%225a9f382a-701a-4ee7-97b4-d9eb45d47c4e%22; ab_testing_id=%22f104b738-5eac-4897-a976-b1f4c2d6b2d8%22; __cf_bm=g3SCauH6z0B2KDLCgcsKV4RELyO9F9R8ItPp3I4vkbA-1698009776-0-AU5x8RyYH4FTN8timiuLcBQ+08bDanYBdbc2uRBNz3A+CSD7T46qJ+fu9IcuG6DLTkw5OrT5QyOdK14hDr1n+1s=; AWSALBTG=XM4atleGAcpMjeE8CmYUXf3k1MjO2R7sxFMoA+MOWi/crPGNtY4ufApqc9tyQR+WEHwbZwBq4zOIkJzP5oXkZt1XjL3VTjZ9225fNolV9eeBOFVzve1KZCO4TULOhMXrk4jNbcDAGq6v1WhzOMLxZveYBQFavJQufsLRjyLceap7; AWSALBTGCORS=XM4atleGAcpMjeE8CmYUXf3k1MjO2R7sxFMoA+MOWi/crPGNtY4ufApqc9tyQR+WEHwbZwBq4zOIkJzP5oXkZt1XjL3VTjZ9225fNolV9eeBOFVzve1KZCO4TULOhMXrk4jNbcDAGq6v1WhzOMLxZveYBQFavJQufsLRjyLceap7; visit_id=%7B%22id%22%3A%226077a65c-c5e2-44b0-b5bf-88497cbc9e7e%22%2C%22timestamp%22%3A%222023-10-22T21%3A23%3A31.512Z%22%7D' \
  -H 'dnt: 1' \
  -H 'referer: https://www.astralcodexten.com/archive?sort=new' \
  -H 'sec-ch-ua: "Google Chrome";v="117", "Not;A=Brand";v="8", "Chromium";v="117"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "macOS"' \
  -H 'sec-fetch-dest: empty' \
  -H 'sec-fetch-mode: cors' \
  -H 'sec-fetch-site: same-origin' \
  -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36' \
  --compressed
```

Get the comments of an article (`123307142` is the article `id`):

```shell
curl 'https://www.astralcodexten.com/api/v1/post/123307142/comments?token=&all_comments=true&sort=oldest_first' \
  -H 'authority: www.astralcodexten.com' \
  -H 'accept: */*' \
  -H 'accept-language: en-GB,en;q=0.9,fr-FR;q=0.8,fr;q=0.7,en-US;q=0.6,ja;q=0.5' \
  -H 'cookie: ajs_anonymous_id=%225a9f382a-701a-4ee7-97b4-d9eb45d47c4e%22; ajs_anonymous_id=%225a9f382a-701a-4ee7-97b4-d9eb45d47c4e%22; ab_testing_id=%22f104b738-5eac-4897-a976-b1f4c2d6b2d8%22; __cf_bm=g3SCauH6z0B2KDLCgcsKV4RELyO9F9R8ItPp3I4vkbA-1698009776-0-AU5x8RyYH4FTN8timiuLcBQ+08bDanYBdbc2uRBNz3A+CSD7T46qJ+fu9IcuG6DLTkw5OrT5QyOdK14hDr1n+1s=; visit_id=%7B%22id%22%3A%226077a65c-c5e2-44b0-b5bf-88497cbc9e7e%22%2C%22timestamp%22%3A%222023-10-22T21%3A23%3A31.512Z%22%7D; AWSALBTG=hCwyt+5N3Sr4QpdawetLpPxBibVbkfYh/Fu25Ltuzn7aX0r2ujLAkKVxlmq0tWUA5AVcDm1OSIsg7FBG7/2SV7McW19tZtkDYyEqsFwofLAoowEeZ8tLlgyPJ65msIQkoOTLcOadmaNnoLTKdIXN5GoH7k207fbQoELRrSqPdvrU; AWSALBTGCORS=hCwyt+5N3Sr4QpdawetLpPxBibVbkfYh/Fu25Ltuzn7aX0r2ujLAkKVxlmq0tWUA5AVcDm1OSIsg7FBG7/2SV7McW19tZtkDYyEqsFwofLAoowEeZ8tLlgyPJ65msIQkoOTLcOadmaNnoLTKdIXN5GoH7k207fbQoELRrSqPdvrU' \
  -H 'dnt: 1' \
  -H 'referer: https://www.astralcodexten.com/p/your-book-review-lying-for-money' \
  -H 'sec-ch-ua: "Google Chrome";v="117", "Not;A=Brand";v="8", "Chromium";v="117"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "macOS"' \
  -H 'sec-fetch-dest: empty' \
  -H 'sec-fetch-mode: cors' \
  -H 'sec-fetch-site: same-origin' \
  -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36' \
  --compressed
```

## Notes

### Old version (writing the JSON, and then inserting the JSON in the database)

- Using transactions with `tx.Stmt` to transform a prepared statement to a transaction-specific prepared statement offered a humongus speedup. SQL went from being ~90% of the runtime (as per pprof) to ~25%
- now the runtime is dominated by `json.Unmarshal`, ~65%

### New version (writing directly in the database)

- Runtime seems to be mostly `time.Sleep(1 * time.Second)` (1 API call every second be respectful)