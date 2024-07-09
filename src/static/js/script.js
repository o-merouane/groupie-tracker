document.addEventListener('click', function(event) {
    const expandedSquare = document.querySelector('.square.expanded');
    if (expandedSquare && !expandedSquare.contains(event.target)) {
        collapseAllSquares();
    }
});

function expandSquare(element) {
    const isExpanded = element.classList.contains('expanded');
    
    if (!isExpanded) {
        // Add 'expanded' class to the clicked square
        element.classList.add('expanded');

        // Get the index of the clicked square
        const gridContainer = document.querySelector('.grid-container');
        const squares = Array.from(gridContainer.children);
        const index = squares.indexOf(element);

        // Calculate the number of columns in the grid
        const gridStyle = window.getComputedStyle(gridContainer);
        const gridColumns = parseInt(gridStyle.getPropertyValue('grid-template-columns').split(' ').length);

        // Calculate the start and end indices of the row
        const rowIndex = Math.floor(index / gridColumns);
        const rowStart = rowIndex * gridColumns;
        const rowEnd = Math.min(rowStart + gridColumns, squares.length);

        // Hide other squares in the same row
        for (let i = 0; i < squares.length; i++) {
            if (i < rowStart || i >= rowEnd) {
                squares[i].classList.add('hidden');
            } else {
                squares[i].classList.remove('hidden');
            }
        }

        // Show more info for the expanded square
        const moreInfo = element.querySelector('.more-info');
        if (moreInfo) {
            moreInfo.style.display = 'flex'; // Show more info
        }
    } else {
        collapseAllSquares(); // Collapse if already expanded
    }
}

function collapseAllSquares() {
    document.querySelectorAll('.square').forEach(square => {
        square.classList.remove('expanded');
        square.classList.remove('hidden');

        const moreInfo = square.querySelector('.more-info');
        if (moreInfo) {
            moreInfo.style.display = 'none'; // Hide more info
        }
    });
}
