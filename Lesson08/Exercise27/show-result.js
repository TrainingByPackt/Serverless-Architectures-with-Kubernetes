function main(params) {
    var examResult = '';
 
    if (params.averageMarks < 0 || params.averageMarks > 100) {
        examResult = 'ERROR: invalid average exam mark';
    } else if (params.averageMarks >= 60) {
        examResult = 'Pass';
    } else {
        examResult = 'Fail';
    }
	
    return { result: examResult };
}
