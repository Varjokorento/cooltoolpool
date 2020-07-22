document.addEventListener('click', function (event) {

    let counter = 0;
	// If the clicked element doesn't have the right selector, bail
	if (!event.target.matches('.click')) return;

	// Don't follow the link
	event.preventDefault();
    counter = counter + 1;
	// Log the clicked element in the console
	console.log(counter);

}, false);