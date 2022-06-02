# go-search

A search engine implemented by golang.

## technique details

|Topic|Detail|Engineer|
-|-|-|
|**tokenizer**|1. Use Jieba to tokenize input sentence.|@Xubaidu|
|**index**|1. Build reverted index for tokens.|@Xubaidu|
|**storage**|1. Use leveldb to store data (need to learn how to use it).||
|**query**|1. Load data from file to memory and write query APIs.<br> 2. Need to evaluate different data structures and algorithms for HPC search.<br>3. Parallely quering levelDB may be considered.||
|**service**|1. Encapsulate algorithm into go service.<br>2. Go gin framework is preferred.||