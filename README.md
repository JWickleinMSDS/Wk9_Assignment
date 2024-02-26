***Changes I made to the original code***

1) Images:  I replaced the four images in the original code with 4 pictures of dragons. (Because dragons are cool.)

2) I added code to perform error checking at several levels.  In the "image_processing.go" file I added error handling logic to the Read and Write image functions. In the "main.go" file I added error handling to the load and save image functions.  The code I provided should not result in any errors.  But, if you want to see the error handling work, simply change the ".jpeg" extension on one of the image files to something else, such as, ".jp" and it should invoke the error handling. 

3) I created a unit test that attempts to compare the pre-coverted image to the post-converted image that was resized to 500 x 500 pixels and made into grayscale.  This is under the "image_processing_test.go" file.  Within the image_processing directory I added two files under the "testdata" directory.  The "inputimage.jpeg" file is the pre-converted, larger color image.  The "expectedGrayImage.jpeg" file is what the converted image should look like.  For some reason, I am unable to pass this unit test.  This is befuddling as I believe that I have done everything right yet the test still failed to pass.  

4) I performed benchmarking of 5 components:  The pipeline (which benchmarks all the steps in the pipeline), loading the images, resizing the images, converting them to grayscale, and saving the images.  The benchmark tests can be found in the file "pipeline_bench_test.go".  Loading images are by far the slowest of the operations.  This is probably the primary area of opportunity when looking to optimize performance.

***Instructions to run the program***

1) Run "go run main.go".
2) Run "go test -v -bench=".""  (Note:  Be sure to keep the quotes around the dot.  Execution of the benchmarks only worked with the quotes - for some reason.)
3) Change the directory to "image_processing" and execute the "go test" command.  This will run the unit tests.  (see note #3 above)



***All of the content below this was part of the original README.md file.***
# Episode #21: Concurrency in Go: Pipeline Pattern

[Episode link](https://www.codeheim.io/courses/Episode-21-Concurrency-in-Go-Pipeline-Pattern-65c3ca14e4b0628a4e002201)

Requires Golang 1.20 or higher.
