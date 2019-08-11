function main(params) {
    var examResult = '';
 
    if (params.examMarkes < 0 || params.examMarkes > 100) {
        examResult = 'ERROR: invalid exam mark';
    } else if (params.examMarkes > 80) {
        examResult = 'Pass with grade A';
    } else if (params.examMarkes > 70) {
        examResult = 'Pass with grade B';
    } else if (params.examMarkes > 60) {
        examResult = 'Pass with grade C';
    } else {
		examResult = 'Fail';
    }
	
    return { result: examResult };
}
