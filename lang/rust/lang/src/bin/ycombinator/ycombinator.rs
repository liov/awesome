//! A simple implementation of the Y Combinator
// λf.(λx.xx)(λx.f(xx))
// <=> λf.(λx.f(xx))(λx.f(xx))

// CREDITS: A better version of the previous code that was posted here, with detailed explanation.
// See <y> and also <y_apply>.

// A function type that takes its own type as an input is an infinite recursive type.
// We introduce a trait that will allow us to have an input with the same type as self, and break the recursion.
// The input is going to be a trait object that implements the desired function in the interface.
// NOTE: We will be coercing a reference to a closure into this trait object.
trait Apply<T, R> {
    fn apply(&self, apply:&dyn Apply<T, R>, t:T ) -> R;
}

// In Rust, closures fall into three kinds: FnOnce, FnMut and Fn.
// FnOnce assumed to be able to be called just once if it is not Clone. It is impossible to
// write recursive FnOnce that is not Clone.
// All FnMut are also FnOnce, although you can call them multiple times, they are not allow to
// have a reference to themselves. So it is also not possible to write recursive FnMut closures
// that is not Clone.
// All Fn are also FnMut, and all closures of Fn are also Clone. However, programmers can create
// Fn objects that are not Clone
// The following address all closures that is Clone, and those are Fn.
impl<T, R, F> Apply<T, R> for F where F: FnOnce( &dyn Apply<T, R>, T ) -> R + Clone {
    fn apply(&self, f: &dyn Apply<T, R>, t: T ) -> R {
        (self.clone())( f, t )

        // If we were to pass in self as f, we get -
        // NOTE: Each letter is an individual symbol.
        // λf.λt.sft
        // => λs.λt.sst [s/f]
        // => λs.ss
    }
}
//This will work for all Fn objects, not just closures
//And it is a little bit more efficient for Fn closures as it do not clone itself.
//However under 1.26 it is not possible to define both. We will
//need to wait for specialization.
//impl<T, R, F> Apply<T, R> for F where F: Fn( &Apply<T, R>, T ) -> R {
//    fn apply( &self, f: &Apply<T, R>, t: T ) -> R {
//        self( f, t )
//}
//Before 1.26 we have some limitations and so we need some workarounds. But now impl Trait is stable and we can
// write the following:
fn y<T,R>(f:impl FnOnce(&dyn Fn(T) -> R, T) -> R + Clone) -> impl FnOnce(T) -> R {
    |t| (|x: &dyn Apply<T, R>, y| x.apply(x, y))
        (&move |x:&dyn Apply<T, R>, y| f(&|z| x.apply(x, z), y), t)

    // NOTE: Each letter is an individual symbol.
    // (λx.(λy.xxy))(λx.(λy.f(λz.xxz)y))t
    // => (λx.xx)(λx.f(xx))t
    // => (Yf)t
}

//Previous version removed as they are just hacks when impl Trait is not available.

fn fac( n: usize ) -> usize {
    let almost_fac = |f: &dyn Fn(usize) -> usize, x| if x == 0 { 1 } else { x * f( x - 1 ) };
    let fac = y( almost_fac );
    fac( n )
}

fn fib( n: usize ) -> usize {
    let almost_fib = |f: &dyn Fn(usize) -> usize, x| if x < 2 { 1 } else { f( x - 2 ) + f( x - 1 ) };
    let fib = y( almost_fib );
    fib( n )
}

fn optimal_fib( n: usize ) -> usize {
    let almost_fib = |f: &dyn Fn((usize, usize, usize)) -> usize, (i0,i1,x)|
        {
            match x {
                0 => i0,
                1 => i1,
                x => f((i1,i0+i1, x-1))
            }
        };
    let fib = |x|y( almost_fib )((1,1,x));
    fib( n )
}

fn main() {
    println!( "{}", fac( 10 ) );
    println!( "{}", fib( 10 ) );
    println!( "{}", optimal_fib( 10 ) );
    let fact = y2(|f, n| {
        if n == 0 {
            1
        } else {
            n * f.call(n - 1)
        }
    });

    for n in 0..=10 {
        println!("{}: {}", n, fact(n));
    }

    // f: Lazy<u64 -> u64> -> u64 -> u64
    fn f(fac: Lazy<'static, Box<dyn FnOnce(u64) -> u64>>) -> Box<dyn FnOnce(u64) -> u64> {
        Box::new(move |n| {
            if n == 0 {
                1
            } else {
                n * fac()(n - 1)
            }
        })
    }
    println!("{}", y3(f)(5));
}

pub struct Rec<'a, T, U>(&'a dyn Fn(Rec<T, U>, T) -> U);

impl<'a, T: 'a, U: 'a> Rec<'a, T, U> {
    pub fn call(&self, x: T) -> U {
        (self.0)(Rec(self.0), x)
    }
}

pub fn y2<T, U>(f: impl Fn(Rec<T, U>, T) -> U) -> impl Fn(T) -> U {
    move |x| f(Rec(&f), x)
}

type Lazy<'a, T> = Box<dyn FnOnce() -> T + 'a>;

// fix: (Lazy<T> -> T) -> T
fn y3<'a, T, F>(f: F) -> T
    where F: Fn(Lazy<'a, T>) -> T + Copy + 'a
{
    f(Box::new(move || y3(f)))
}
