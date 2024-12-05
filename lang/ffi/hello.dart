import 'dart:ffi' as ffi;


typedef CurrentTimestampFunc = ffi.Int64 Function();
typedef HelloWorld = int Function();

typedef PlusFunc = ffi.Int32 Function(ffi.Int32,ffi.Int32);
typedef Plus = int Function(int,int);

typedef PlusOneFunc = ffi.Int32 Function(ffi.Int32);
typedef PlusOne = int Function(int);

final lib = ffi.DynamicLibrary.open('libnewplus.dll');
final HelloWorld current_timestamp =
lib.lookup<ffi.NativeFunction<CurrentTimestampFunc>>("current_timestamp")
    .asFunction();
final Plus plus =
lib.lookup<ffi.NativeFunction<PlusFunc>>("plus")
    .asFunction();
final PlusOne plusone =
lib.lookup<ffi.NativeFunction<PlusOneFunc>>("plusone")
    .asFunction();


void run(int count) {
  // start
  var start = current_timestamp();

  int x = 0;
  while (x < count) {
    x = plusone(x);
  }

  print(current_timestamp() - start);
}


void main(List<String> args) {
  if (args.length == 0) {
      print("First arg (0 - 2000000000) is required.");
      return;
  }
  var count = int.parse(args[0]);
  if (count <= 0 || count > 2000000000) {
    print("Must be a positive number not exceeding 2 billion.");
    return;
  }
  plusone(current_timestamp() == 0 ? 1 : 2);
  run(count);

}