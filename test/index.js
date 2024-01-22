var filterObj = {
	itemId: "ItemCode",
	itemName: "ItemName",
	reportDetails: "ItemResult",
	reportAbnormalSign: "AbnormalFlag",
	remark: "Reference",
	itemUnit: "Unit",
};
// 写一个函数根据filterObj的配置，将data中的数据进行过滤，返回一个新的对象
// Path: src/test/st.js
function filterData(data, filterObj) {
	var result = {};
	for (var key in filterObj) {
		result[key] = data[filterObj[key]];
	}
	return result;
}
function defaultAction(data) {
	var res = data.response2DataRow.data;
	var result = [];
	for (var i = 0; i < res.length; i++) {
		var obj = filterData(res[i], filterObj);
		obj["reportTime"] = data.report.reportDate;
		obj["reportDate"] = data.report.reportDate.substring(0, 10);
		result.push(obj);
	}
	return JSON.stringify(result);
}
