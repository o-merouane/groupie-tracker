document.addEventListener('click', function(event) {
    const expandedSquare = document.querySelector('.square.expanded');
    if (expandedSquare && !expandedSquare.contains(event.target)) {
        collapseAllSquares();
    }
});

function expandSquare(element) {
    const isExpanded = element.classList.contains('expanded');
    
    collapseAllSquares();

    if (!isExpanded) {
        // Add 'expanded' class to the clicked square
        element.classList.add('expanded');

        // Hide other squares in the same row
        const gridContainer = document.querySelector('.grid-container');
        const squares = Array.from(gridContainer.children);
        const index = squares.indexOf(element);
        const rowStart = Math.floor(index / 5) * 5;

        for (let i = rowStart; i < rowStart + 5; i++) {
            if (squares[i] !== element) {
                squares[i].classList.add('hidden');
            }
        }
    }
}

function collapseAllSquares() {
    document.querySelectorAll('.square').forEach(square => {
        square.classList.remove('expanded');
        square.classList.remove('hidden');
    });
}
