export const BtnTargetBlur = (e) => {
	// 添加失去焦点事件
	let target = e.target;
	//console.log(e)
	if (target.nodeName === "BUTTON" || target.nodeName === "SPAN") {
		target.parentNode.blur();
	} else if (target.nodeName === "I" || target.nodeName === "svg") {
		target.parentNode.parentNode.blur();
	} else if (target.nodeName === "svg" || target.nodeName === "path") {
		target.parentNode.parentNode.parentNode.blur();
	} else if (target.nodeName === "path") {
		target.parentNode.parentNode.parentNode.parentNode.blur();
	}
	target.blur();
}
