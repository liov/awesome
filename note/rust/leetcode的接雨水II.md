与接雨水I难度同样是困难，但比I难多了，我用了两个版本，原因是因为二叉树频繁创建迭代器让我不满

两个版本测试对比，800ms的是二叉树的

提交时间	状态	执行用时	内存消耗	语言
几秒前	通过	776 ms	3 MB	rust
4 小时前	通过	800 ms	3 MB	rust

版本一：二叉树版本
```rust
use std::collections::{HashSet, BTreeMap};


#[derive(PartialEq, Eq, PartialOrd, Ord, Debug, Hash, Clone, Copy)]
struct Point(i32, usize, usize);


pub fn trap_rain_water(height_map: Vec<Vec<i32>>) -> i32 {
    if height_map.len() < 3 || height_map[0].len() < 3 { return 0; };
    let m = height_map.len();
    let n = height_map[0].len();
    let mut result = 0;
    let mut side = BTreeMap::new();
    let mut drop = HashSet::new();
    for x in 0..m {
        for y in 0..n {
            if x == 0 || x == m - 1 || y == 0 || y == n - 1 {
                side.insert(Point(height_map[x][y], x, y), false);
                drop.insert((x, y));
            }
        }
    }

    let round: [[i32; 2]; 4] = [[0, -1], [0, 1], [-1, 0], [1, 0]];
    let mut x = 0;
    let mut y = 0;

    let mut point_iter = side.iter();
    while let Some((point, yet)) = point_iter.next() {
        if *yet { continue; };
        let point = *point;
        for j in round.iter() {
            x = (point.1 as i32 + j[0]) as usize;
            y = (point.2 as i32 + j[1]) as usize;
            if x >= m - 1 || y >= n - 1 {
                continue;
            }
            if drop.get(&(x, y)) != None {
                continue;
            }
            drop.insert((x, y));
            if height_map[x][y] <= point.0 {
                result = result + (point.0 - height_map[x][y]);
                side.insert(Point(point.0, x, y), false);
            } else {
                side.insert(Point(height_map[x][y], x, y), false);
            }
            point_iter = side.iter();
        }
        side.insert(point, true);
        point_iter = side.iter();
    }
    result
}
```
版本二：纯数组做插入排序
```rust
//用Vec插入排开销太大，放弃
use std::fmt::{Display, Debug};
#[derive(PartialEq, Eq, PartialOrd, Ord, Debug, Hash, Clone, Copy)]
struct Point2(i32, usize, usize, bool);


pub fn insert_sort<T>(ord_vec:&mut Vec<T>, mut other:Vec<T>) where T: Ord+Copy+Debug {
    other.sort_by(|a, b| b.cmp(a));
    let len = ord_vec.len();
    for i in 0..other.len() {
        ord_vec.push(other[i]);
    }

    if other[0]<=ord_vec[len-1]{return;}

    for i in (0..len).rev() {
        let other_len =other.len();
        if other_len == 0 { break; }
        for j in 0..other_len {
            if other[j] <= ord_vec[i] {
                ord_vec[i + j] = ord_vec[i];
                for x in (j..other_len).rev() {
                    ord_vec[i + x + 1] = other[x];
                    other.pop();
                }
                break;
            }
        }
        if other.len() == other_len{
            ord_vec[i+other_len]=ord_vec[i];
        }
    }
    if other.len()>0{
        for j in 0..other.len(){
            ord_vec[j+other.len()]=ord_vec[j];
            ord_vec[j]=other[j];
        }
    }
}


pub fn trap_rain_water2(height_map: Vec<Vec<i32>>) -> i32 {
    if height_map.len() < 3 || height_map[0].len() < 3 { return 0; };
    let m = height_map.len();
    let n = height_map[0].len();
    let mut result = 0;
    let mut side = Vec::with_capacity(m * n);
    let mut drop = HashSet::new();
    for x in 0..m {
        for y in 0..n {
            if x == 0 || x == m - 1 || y == 0 || y == n - 1 {
                side.push(Point2(height_map[x][y], x, y, false));
                drop.insert((x, y));
            }
        }
    }
    side.sort_by(|a, b| b.cmp(a));
    let round: [[i32; 2]; 4] = [[0, -1], [0, 1], [-1, 0], [1, 0]];
    let mut x = 0;
    let mut y = 0;
    let mut i = side.len() - 1;
    loop {
        if side[i].3 {
            i = i - 1;
            continue;
        }
        let mut sub_side = Vec::with_capacity(3);

        for j in round.iter() {
            x = (side[i].1 as i32 + j[0]) as usize;
            y = (side[i].2 as i32 + j[1]) as usize;
            if x < 0 || x >= m - 1 || y < 0 || y >= n - 1 {
                continue;
            }
            if drop.get(&(x, y)) != None {
                continue;
            }
            drop.insert((x, y));
            if height_map[x][y] <= side[i].0 {
                result = result + (side[i].0 - height_map[x][y]);
                sub_side.push(Point2(side[i].0, x, y, false));
            } else {
                sub_side.push(Point2(height_map[x][y], x, y, false))
            }
            //这里应该有个else，插入排大于边的新边
        }
        side[i].3 = true;
        if sub_side.len() > 0 {
            insert_sort(&mut side,sub_side);
            i = side.len();
        }

        if i == 0 { break; }
        i = i - 1;
    }
    result
}
```
二比一代码量多了很多，开发耗时自然多，重要原因是自己造插入排序的轮子，不像二叉树拿来即用

意外的纯数组耗时居然比二叉树少，虽然不少多少

