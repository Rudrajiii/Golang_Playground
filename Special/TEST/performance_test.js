const ffi = require('ffi-napi');
const ref = require('ref-napi');


const lib = ffi.Library('./libgo_function.so', {
  'GoFunction': ['void', []]
});

// JavaScript function to compare
function jsFunction() {
  let sum = 0;
  for (let i = 0; i < 1000000000; i++) {
    sum += i;
  }
  console.log("JS sum:", sum); // Prevent optimization
}

// Function to measure execution time
function measureExecutionTime(fn, label) {
  const start = process.hrtime.bigint();
  fn();
  const end = process.hrtime.bigint();
  const duration = Number(end - start) / 1e6; // Convert nanoseconds to milliseconds
  console.log(`${label} took ${duration.toFixed(3)} ms`);
  return duration;
}

// Performance test for Go and JavaScript functions
function runPerformanceTest(iterations = 10) {
  console.log(`Starting performance test (${iterations} iterations each)...`);

  let goTotalTime = 0;
  let jsTotalTime = 0;

  for (let i = 0; i < iterations; i++) {
    // Measure Go function performance
    const goDuration = measureExecutionTime(() => lib.GoFunction(), `Go (iteration ${i + 1})`);
    goTotalTime += goDuration;

    // Measure JavaScript function performance
    const jsDuration = measureExecutionTime(jsFunction, `JavaScript (iteration ${i + 1})`);
    jsTotalTime += jsDuration;
  }

  // Calculate and display average times
  const goAvg = goTotalTime / iterations;
  const jsAvg = jsTotalTime / iterations;

  console.log("\nAverage Times:");
  console.log(`Go: ${goAvg.toFixed(3)} ms`);
  console.log(`JavaScript: ${jsAvg.toFixed(3)} ms`);

  // Compare performance
  const speedup = jsAvg / goAvg;
  console.log(`\nGo is approximately ${speedup.toFixed(2)}x faster than JavaScript.`);
}

runPerformanceTest(10);
