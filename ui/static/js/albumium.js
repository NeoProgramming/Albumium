function openURL(event) {
    event.preventDefault();
    let appId = document.getElementById("app_id").value;
    console.log("openURL: " + appId);
    let appUrl = "https://oauth.vk.com/authorize?client_id="
        + appId + "&display=page&redirect_uri=https://oauth.vk.com/blank.html&scope=notify,friends,photos,audio,video,docs,notes,pages,status,wall,groups,notifications&response_type=token&v=5.131"
    // open in new page
    window.open(appUrl, "_blank");
    // prepare AJAX
    let xhr = new XMLHttpRequest();
    xhr.open("POST", "/set-app-id", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    // handler
    xhr.onreadystatechange = function() {
        if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            //alert(xhr.responseText);
            location.reload()
        }
    };

    xhr.send("app_id=" + appId);
}

function postURL(event) {
    event.preventDefault();
    let appUrl = document.getElementById("app_url").value;
    console.log("postURL: " + appUrl);

    // window.open("https://google.com", "_blank");
    // prepare AJAX
    let xhr = new XMLHttpRequest();
    xhr.open("POST", "/set-app-url", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    // handler
    xhr.onreadystatechange = function() {
        if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            //alert(xhr.responseText);
            location.reload()
        }
    };
    // send
    let prm = "app_url=" + appUrl;
    console.log("send to: " + prm);
    xhr.send(prm);
}

function setProxy(event) {
    event.preventDefault();
    let proxyUrl = document.getElementById("proxy_addr").value;
    let proxyUse = document.getElementById("proxy_use").checked;
    console.log("proxyURL: " + proxyUrl + " , proxyUse: " + proxyUse);
    let xhr = new XMLHttpRequest();
    xhr.open("POST", "/set-proxy", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    // handler
    xhr.onreadystatechange = function() {
        if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            //alert(xhr.responseText);
            location.reload()
        }
    };
    // send
    let prm = "proxy_url=" + proxyUrl + "&proxy_use=" + proxyUse;
    console.log("proxy prm: " + prm);
    xhr.send(prm);
}

function openTestURL(ibase) {
    let base = document.getElementById(ibase).textContent;
    console.log(base)
    let token = document.getElementById('token').textContent;
    console.log(token)
    let url = base + token;
    window.open(url, "_blank");
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
    let currentUrl = window.location.href.split('?')[0];
    let code = getChk('f_my') +  getChk('f_bm') + getChk('f_fr') + getChk('f_gr') + getChk('f_lk') + getChk('f_cm');
    if(filtersIsEmpty(code))
		window.location.href = currentUrl + '?' +extraArgs;
	else
		window.location.href = currentUrl + '?' +pkArgs('filters', code, extraArgs);
}

//
function openById(vkurl) {
    // extract type and id, redirect to local page
    let url = new URL(vkurl);
    let path = url.pathname;
    let parts = path.match(/([a-zA-Z]+)([0-9]+)/);
    let type = parts[1]; // "id"
    let number = parts[2]; // "12345678"
    if(type=="id") {

    } else if(type=="public") {

    } else {

    }
}

// HELPERS FOR UPDATE DB QUERIES

function sendUpdateQuery(url) {
    // prepare AJAX
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
    // prepare AJAX
    let xhr = new XMLHttpRequest();
    xhr.open("POST", url, true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");

    const data = "id=" + argvalue;

    console.log(url);
    xhr.send(data);
}

function sendUpdateQueryCB(url) {
    // prepare AJAX
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

// GLOBAL AREA

document.addEventListener("DOMContentLoaded", function(event) {
    console.log("init page")

    let statusbar = document.getElementById('statusbar');

    // Update the status in the HTML
    let source = new EventSource("http://localhost:8080/get-server-status");
    source.onmessage = function(event) {
        statusbar.textContent = event.data;
    };

    //
    let started = false;
    let btn = document.getElementById("controlbtn");

    // Get the initial state of the goroutine
    let xhr = new XMLHttpRequest();
    xhr.open("GET", "/get-worker-status");
    xhr.onload = function() {
        if (xhr.status === 200) {
            started = (xhr.responseText === "true");
            btn.textContent = started ? "STOP" : "START";
        }
    };
    xhr.send();

    btn.addEventListener("click", function(event) {
        event.preventDefault();
        if (started) {
            console.log("stop...")
            btn.textContent = "STOPPING...";
            // Stop the goroutine
            let xhr = new XMLHttpRequest();
            xhr.open("POST", "http://127.0.0.1:8080/stop-worker");
            xhr.onload = function() {
                console.log("Worker stopped.");
                started = false;
                btn.textContent = "START";
            };
            xhr.onerror = function() {
                console.log("Error stopping worker.");
            };
            xhr.send();
        } else {
            console.log("start...")
            btn.textContent = "STARTING...";
            // Start the goroutine
            let xhr = new XMLHttpRequest();
            xhr.open("POST", "http://127.0.0.1:8080/start-worker");
            xhr.onload = function() {
                console.log("Worker started.");
                started = true;
                btn.textContent = "STOP";
            };
            xhr.onerror = function() {
                console.log("Error starting worker.");
            };
            xhr.send();
        }
    });
});

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
