function main(params) {
    var examResult = '';
 
    if (params.examMarks < 0 || params.examMarks > 100) {
        examResult = 'ERROR: invalid exam mark';
    } else if (params.examMarks > 80) {
        examResult = 'Pass with grade A';
    } else if (params.examMarks > 70) {
        examResult = 'Pass with grade B';
    } else if (params.examMarks > 60) {
        examResult = 'Pass with grade C';
    } else {
		examResult = 'Fail';
    }
	
    return { result: examResult };
}
