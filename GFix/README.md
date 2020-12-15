# GFix: automatic patching tool of goroutine leak bugs

## Examples and demonstration

In `examples/src` directory, there are examples for GL-1, GL-2 and GL-3 respectively. `examples/run.sh` makes a demonstration of how to use the tools on them.

## Directories

1. `dispatcher`: The tool to determine if the program could be patched by GL-1, GL-2, GL-3 or none. How to use could be found in `examples/run.sh`. 
Parameters: `-buggyfilepath=/absolute/path/to/the/buggyfile.go` `-path=/path/to/the/go/project` `-makelineno=line_number_of_the_make_channel_operation` `-oplineno=line_number_of_the_blocking_operation` `-include=/other/necessary/paths/to/compile/the/project`
The output starting with `[DISPATCH]` indicates the program could be patched by GL-1 to GL-3 if it is 1, 2 or 3. Zero indicates it could not be patched. `[PATCH]` indicates the line number(s) for the patcher.

2. `gl-1-patcher`: The patcher for GL-1. Usage: `./bin/gl1_patch /path/to/buggy_file.go line_of_make_chan_to_patch`.
3. `gl-2-patcher`: The patcher for GL-2. Usage: `./bin/gl2_patch /path/to/buggy_file.go line_to_insert_defer line_1_to_remove line_2_to_remove ...`.
4. `gl-3-patcher`: The patcher for GL-3. Usage: `./bin/gl3_patch /path/to/buggy_file.go line_to_insert_the_new_channel first_line_of_channel_operation_to_change second_line_of_channel_operation_to_change ...`.
