#!/bin/bash

# Change imports and go:generate strings to new path.
grep -rl github.com\/google\/wire . | xargs sed -i 's/github.com\/google\/wire/gortc.io\/wire/g'
grep -rl '\/\/go:generate wire' . | xargs sed -i 's/\/\/go:generate wire/\/\/go:generate go run gortc.io\/wire\/cmd\/wire/g'

# Change github links back.
grep -rl 'https:\/\/gortc.io\/' . | xargs sed -i 's/https:\/\/gortc.io\//https:\/\/github.com\/google\//g'