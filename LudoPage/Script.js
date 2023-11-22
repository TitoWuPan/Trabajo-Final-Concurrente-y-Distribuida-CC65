const startButton = document.getElementById('startButton');
const textContainer = document.querySelector('.text-container');
const mazeContainer = document.getElementById('mazeContainer');

startButton.addEventListener('click', function () {
    const newTask = document.createElement('p');
    
    newTask.innerHTML = '<span style="color: #4caf50;"></span> Nueva tarea agregada';

    textContainer.appendChild(newTask);
    textContainer.scrollTop = textContainer.scrollHeight;
});

function adjustTextContainerHeight() {
    var windowHeight = window.innerHeight;
    document.getElementById('dynamicTextContainer').style.maxHeight = windowHeight + 'px';
}

function createMazeTable() {
    const matrix = [
        [1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1],
        [1, 3, 0, 0, 0, 2, 0, 0, 0, 0, 0, 1],
        [1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1],
        [1, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 1],
        [1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1],
        [1, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 1],
        [1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1],
        [1, 0, 1, 0, 0, 0, 1, 0, 2, 0, 0, 1],
        [1, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1],
        [1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 0, 1],
        [1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 4, 1],
        [1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1],
    ];

    const table = document.createElement('table');

    for (let i = 0; i < matrix.length; i++) {
        const row = document.createElement('tr');

        for (let j = 0; j < matrix[i].length; j++) {
            const cell = document.createElement('td');

            if (matrix[i][j] === 1) {
                cell.className = 'wall';
            } else if (matrix[i][j] === 2) {
                cell.className = 'trap';
            } else if (matrix[i][j] === 3) {
                cell.className = 'start';
            } else if (matrix[i][j] === 4) {
                cell.className = 'end';
            } else {
                cell.className = 'empty';
            }
            row.appendChild(cell);
        }
        table.appendChild(row);
    }
    mazeContainer.appendChild(table);
}


