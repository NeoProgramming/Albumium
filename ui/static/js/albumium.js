function openMedia(id) {
   var xhr = new XMLHttpRequest();
   xhr.open("POST", "/open-media/" + id, true);
   console.log("open-media " , id);
   xhr.send();
}

function onCheckMedia(id) {
	console.log("onCheckMedia " , id);
	var checkbox = document.getElementById(id);
    var parentDiv = checkbox.parentNode.parentNode;
    if(checkbox.checked) {
       parentDiv.style.backgroundColor = 'red';
    } else {
       parentDiv.style.backgroundColor = 'white';
    }
}

function setSearch(extraArgs) {
	console.log("setSearch ", extraArgs);
    // Get the name from the form
    let text = document.getElementById('search').value;
    let currentUrl = window.location.href.split('?')[0];
    window.location.href = currentUrl + '?' +pkArgs('search', encodeURIComponent(text), extraArgs);
}

function clearSearch(extraArgs) {
    let currentUrl = window.location.href.split('?')[0];
	window.location.href = currentUrl + '?' + extraArgs;
}

function applyFilters(extraArgs) {
	console.log("applyFilters ", extraArgs);
    let currentUrl = window.location.href.split('?')[0];
    let code = getChk('f_photo') +  getChk('f_video');
    if(filtersIsEmpty(code))
		window.location.href = currentUrl + '?' +extraArgs;
	else
		window.location.href = currentUrl + '?' +pkArgs('filters', code, extraArgs);
}

function changeFileTypes() {
	
}

// HELPERS FOR UPDATE DB QUERIES

function sendUpdateQuery(url) {

    let xhr = new XMLHttpRequest();
    xhr.open("POST", url, true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    // handler
    xhr.onreadystatechange = function() {
        if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            //alert(xhr.responseText);
            //location.reload()
            // while the handler function is unnecessary?
        }
    };

    console.log(url);
    xhr.send();
}

function sendUpdateQueryARG(url, argvalue) {
    console.log("sendUpdateQueryARG: ", argvalue)

    let xhr = new XMLHttpRequest();
    xhr.open("POST", url, true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");

    const data = "id=" + argvalue;

    console.log(url);
    xhr.send(data);
}

function sendUpdateQueryCB(url) {
    let xhr = new XMLHttpRequest();
    xhr.open("POST", url, true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");

    const checkboxes = document.querySelectorAll("input[type=checkbox]:checked");
    const checkboxValues = [];
    for (let i = 0; i < checkboxes.length; i++) {
        if(checkboxes[i].id != "all")
            checkboxValues.push(checkboxes[i].id);
    }
    const data = "checkbox=" + encodeURIComponent(checkboxValues.join(","));

    xhr.send(data);
    console.log(url);
 }

function checkAll()
{
    const mainCheckbox = document.getElementById('all');
    const checkboxes = document.querySelectorAll('input[type="checkbox"]');

    mainCheckbox.addEventListener('click', function() {
        checkboxes.forEach(function(checkbox) {
            checkbox.checked = mainCheckbox.checked;
        });
    });
}

function ts(cb) {
    if (cb.readOnly) cb.checked=cb.readOnly=false;
    else if (!cb.checked) cb.readOnly=cb.indeterminate=true;
}

function getChk(id) {
    let cb = document.getElementById(id);
    if(cb.indeterminate) return '2';
    if(cb.checked) return '1';
    return '0';
}

function setChk(id, st) {
    let cb = document.getElementById(id);
    if(st=='2') {
		cb.readOnly=cb.indeterminate=true;
	} else if(st=='1') {
		cb.readOnly=false;
		cb.checked=true;
	} else if(st=='0') {
		cb.checked=cb.readOnly=false;
	}
}

function filtersIsEmpty(str) {
	for (let i = 0; i<str.length; i++)
		if(str[i]!='0')
			return false;
	return true;
}

function pkArgs() {
    let i = 0;
	let res = '';
    // pairs
	for (i = 0; i < arguments.length-1; i+=2) {
		if(arguments[i+1] != '') {
			if(i>0)
				res += '&';
			res += arguments[i];
            res += '=';
            res += arguments[i + 1];
		}
    }
    // trailing argument
    if(i < arguments.length) {
        if(i>0)
            res += '&';
        res += arguments[i];
    }
	return res;
}

// GLOBAL AREA

document.addEventListener("DOMContentLoaded", function(event) {
    console.log("init page")
    
});


