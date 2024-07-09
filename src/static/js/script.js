function expandSquare(element) {
    const isExpanded = element.classList.contains('expanded');
    
    // Remove 'expanded' class from all squares
    document.querySelectorAll('.square').forEach(square => {
        square.classList.remove('expanded');
        square.classList.remove('first');
    });

    if (!isExpanded) {
        // Add 'expanded' class to the clicked square
        element.classList.add('expanded');

        // Check if the square is the first in the row
        const gridContainer = document.querySelector('.grid-container');
        const squares = Array.from(gridContainer.children);
        const index = squares.indexOf(element);
        const isFirst = index % 5 === 0;

        if (isFirst) {
            element.classList.add('first');
        }
    }
}
