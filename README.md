# A quick throughput test

I saw [this](https://lemire.me/blog/2021/08/03/how-fast-can-you-pipe-a-large-file-to-a-c-program/) post about pipe throughput on an M1 Mac vs. an ARM-based Centos machine and was intrigued. I wrote the fastest, dumbest emitter and consumer I could to try to test it out. The buffers used here are enormous (4MB) to the point of being silly, but the goal was to test the upper limit.

## Results

I ran the test on my Macbook Pro (13" 2020, 2.3GHz Intel i7, 32GB RAM) and made no attempts to make a clean environment -- I even left Chrome running. I ran the test two ways:

### Straight

```bash
go run . emit | go run .
total:        34359738368
duration:     7.148405181s
rate (GiBps): 4.477
```

### Through [pv](https://linux.die.net/man/1/pv)

```bash
go run . emit | pv | go run .
32.0GiB 0:00:08 [3.76GiB/s]
total:        34359738368
duration:     8.209214543s
rate (GiBps): 3.898
```

I'd say these results are within the margin of error of each other. These are my conclusions:

* The 0.04GBps that the original author reported are far below the potential of the system
* Standard tools (such as `pv`) are capable of efficiently handling piped data on the Mac

Conclusions I can't draw:

* This test tells nothing about differences between M1 (ARM) and Intel performance -- I don't have an M1 to test with
* Smaller and/or mismatched buffer sizes could change the result significantly
