## Why make benchmark for skiplist

I create a skip list repo after reading the paper, "skip lists： A probabilistaic alternative to balanced treees.pdf". then I find a [skiplist repo survey](https://github.com/sean-public/skiplist-survey), which is lack of maintain.

So I add and refactor a trunk of code from the survey，   then created this repo. Of cource, i fixed some old summries and review code from other guys, which benefit a lot. Finnally, I will complete the benchmark result in the file and optimize my [skiplist](https://github.com/Ruth-Seven/skiplist)


## Benchmark of Skip List Implementations

Here is a brief summary of skip list packages available in Go that you may consider using after a quick Google/Github search. If you know of any others, please contact me so I can add them here.

Some things most of the packages have in common:

- Without generics in go, Using **interface** to fill key and value in elements. And you must provided `compare()` to determine the order of all the elements.
- The probability of adding new nodes to each linked level is *P*. The values vary from 0.25 to 0.5.In almost of implementations, p can reset dynamical. This is an important parameter for performance tuning and memory usage. You can figure out reasons in [there](https://15721.courses.cs.cmu.edu/spring2018/papers/08-oltpindexes1/pugh-skiplists-cacm1990.pdf).
- Providing some features, such as  `Iterator` 、 `Insert()`、 `Delete()`、`Find()` 、`Element Create`

### Brief Note on each implementation

- [mtchavez-skiplist](https://github.com/mtchavez/skiplist)
  - thread-safe but using coarse-grained lock to prevent race
  - normal performace
- [huandu-skiplist](https://github.com/huandu/skiplist)
  - Globally sets *P* to *almost* 0.25 (using bitmasks and shifting) and can be changed at runtime.
  - Not threadsafe while only uising `mutex` in `list.insert()` for unknown reasons📌
  - Normal performance
- [zhenjl-skiplist](https://github.com/zhenjl/skiplist)
  - slowest💊
  - Adjustable *P* value and max level per list.
  - Allows **duplicates** stored at a single key and therefore does not have an update operation.
  - Uses separate search and insert fingers to speed up finding highly local keys consecutively.
  - Threadsafe but unscalability. The fingers are shared as well across all lists
- [go-datastructures-skip](https://github.com/golang-collections/go-datastructures/slice/skip)
  - very very **slow**⭕️
  - *Over-designed* SkipList Node using a lot of contracts.
  - You can find element by **index**, but it make skiplist more complex and more slow.
  - Provide `Split()` 、`Replace` and `Iterator` API.
- [ryszard-goskiplist](https://github.com/ryszard/goskiplist)
  - P value is a global constant, 0.25
  - ⭕️**Very slow** becase of using a lof of interface to save and compare, even with comparision function. Maybe he will like **genenic** in go2.
  - Provide `Map` and `Set` data structure API
- [sean-skiplist](https://github.com/sean-public/fast-skiplist)
  - Sample and comcise, more IMPORTANTLY, **faster✅**
- [MauriceGit-skiplist](https://github.com/MauriceGit/skiplist)
  - [Delete Data in the png](https://github.com/MauriceGit/skiplist) is too slow. And I found he make a [litter trick](https://github.com/MauriceGit/skiplist-survey/blob/master/mtchavez.go) to bypass the obstacles of benchmark. I don't think it's fair.
  - Using a different algorithm to find element in list. And it's a little complex.
  - Using tarilzeros of genered number to generate random level. It's **creative** and means a the fixed p.
  - Implement of `list.String()` is very help for debuging list and show. But it also request that element must implement `String()` interface. 
  - Of course, the skiplist is very quick!✅

### Run

Running the benchmarks found in this repo locally is easy:

2. run the commands

```sh
git clone https://github.com/Ruth-Seven/skiplists_benchmark.git && cd ./skiplists_benchmark 
pip install pathlib pandas numpy seaborn 
./run.sh 10 10 1000 example  
# the parament means as follows:
# Iterator start 
# factor of iterator multipile
# Iterator end
# output folder
```

2. check updated result pngs in [Visualization.md](./example/png.md) contained by this file.

The results are in CSV format for easy charting and analysis.

### Terms explaination

- For each char in pngs, the vertical axis is **nanoseconds per operation**, the horizontal is the number of items in the list.
- **Best-case insert(insert)**: These are the "best" inserts because they happen at the front of the list, which shouldn't require any searching.
- **Worst-case inserts**: These inserts are at the end of the list, requiring searching all the way to the end and then adding the new node. 
- **Random inserts**: The inserts are at random positions in the skiplist, making this the closest real-world case for inserts. The approximately logarithmic behaviour is clearly visible for all implementations.
- **WorstSearch**: like **Worst-case inserts.** 
- **Average search deletions**: Sequentailly delete element in a skiplist.
- **Best-case Deletes(delete)**: like best-case insert.
- **Worst case deletions**:  In this benchmark, a skip list of a given length is created and then every item is removed, starting from the last one and moving to the front.
- **Random deletions**: Elements are removed from random positions in the skiplist. For Deletions, this is the closest to a real-world case. We can clearly see the the logarithmic behaviour, even though some implementations have a large overhead involved.



### Comparasion

STOP READING BELOW =_=!

PLEASE CHECK DATA BY YOURSELF

THIS IS BULLSHIT>_<

I AM WRITTING NOW.

> 
>
> Even though all implementations only show a small variance even after millions of nodes are added, we can still see very large differences in overall speed because of implementation overhead.
>
> 
>
> `mtchavez`, `sean` and `mauriceGit (mt)` are approximately equally fast. `zhenjl` seems to introduce some serious overhead, making it more than 8x as slow as the fastest implementations.
>
> 
>
> `mauriceGit (mt)` and `sean` are around equally fast with around 100ns faster than the next contestant (`huandu`).
>
> Just like for randomInserts, `mauriceGit (mt)` is the fastest, closely followed by `sean`.

