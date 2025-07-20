# Memory Limit Feature

## Key Features Implemented

### 1. **Memory Limit Setting**
- **64MB limit**: Set a configurable memory limit (currently 64MB)
- **Platform-specific**: Uses `ulimit -v` on Linux/macOS for hard memory limits
- **Cross-platform**: Fallback for other operating systems

### 2. **Real-time Memory Monitoring**
- **Continuous monitoring**: Checks memory usage every 100ms during execution
- **Peak tracking**: Records the maximum memory usage during execution
- **Automatic termination**: Kills the process if memory limit is exceeded

### 3. **Memory Usage Measurement**
- **Linux**: Reads from `/proc/[pid]/status` for accurate RSS memory usage
- **macOS**: Uses `ps` command to get RSS memory usage
- **Cross-platform**: Returns 0 for unsupported platforms

### 4. **Enhanced Error Handling**
- **Memory exceeded detection**: Sets `IsMemoryExceeded` flag when limit is hit
- **Detailed error messages**: Shows the memory limit in error message
- **Accurate reporting**: Reports actual peak memory usage in `MemoryUsageKB`

## How It Works

### 1. **Memory Limit Enforcement**
- Uses `ulimit -v` to set virtual memory limit at OS level
- Monitors process memory in real-time using a goroutine
- Immediately terminates processes that exceed the configured limit

### 2. **Memory Monitoring Process**
- Checks every 100ms for current memory usage
- Tracks peak memory usage throughout execution
- Immediately kills process if limit exceeded
- Uses platform-specific methods for accurate measurement

### 3. **Platform-Specific Implementation**
- **Linux**: Reads RSS memory from `/proc/[pid]/status`
- **macOS**: Uses `ps -o rss=` command for memory measurement
- **Other platforms**: Graceful fallback with basic monitoring

### 4. **Result Reporting**
- `IsMemoryExceeded`: Boolean flag indicating if memory limit was hit
- `MemoryUsageKB`: Reports actual peak memory usage in kilobytes
- `Error`: Contains descriptive error message when limit is exceeded

## Configuration

### Memory Limit Adjustment
You can easily adjust the memory limit by changing the `memoryLimitMB` constant:

```go
const memoryLimitMB = 128 // 128MB memory limit
const memoryLimitMB = 32  // 32MB memory limit
```

### Monitoring Frequency
The monitoring frequency can be adjusted by changing the monitoring interval:

```go
case <-time.After(100 * time.Millisecond): // Check every 100ms
```

## Benefits

1. **Resource Protection**: Prevents memory-hungry code from consuming excessive system resources
2. **Accurate Reporting**: Provides precise memory usage statistics
3. **Automatic Cleanup**: Ensures processes are terminated before they can impact system stability
4. **Cross-platform Support**: Works reliably across different operating systems
5. **Real-time Monitoring**: Immediate detection and response to memory violations

## Technical Implementation

The feature uses a combination of:
- OS-level resource limits (`ulimit`)
- Real-time process monitoring
- Goroutine-based concurrent memory tracking
- Platform-specific memory measurement APIs

This implementation provides robust memory management while maintaining high performance and accuracy in resource usage reporting.
