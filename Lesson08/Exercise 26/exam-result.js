function main(params) {
    var examResult = '';
 
    if (params.examMarks < 0 || params.examMarks > 100) {
        examResult = 'ERROR: invalid exam mark';
    } else if (params.examMarks >= 60) {
        examResult = 'Pass';
    } else {
	  examResult = 'Fail';
    }
	
    return { result: examResult };
}
