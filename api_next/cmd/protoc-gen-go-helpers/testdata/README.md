The classification corpus mirrors `api/cmd/generate-stable-protos/testdata/classification/`.
Keep inputs and expected classification JSON identical in both repositories when
adding a shared schema rule. The copy here lets API-Go CI run independently.

Expected projection text in the API repository is independently authored; never
replace it with actual generator output automatically.
