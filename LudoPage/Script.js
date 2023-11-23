const startButton = document.getElementById('startButton');
const textContainer = document.querySelector('.text-container');
const mazeContainer = document.getElementById('mazeContainer');

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

let valor = 1;

const matrix_aux = [
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

startButton.addEventListener('click', function () {
    const newTask = document.createElement('p');
    
    newTask.innerHTML = '<span style="color: #4caf50;"></span> Nueva tarea agregada';

    MazeTable()

    // insertFicha(1, valor);

    textContainer.appendChild(newTask);
    
    textContainer.scrollTop = textContainer.scrollHeight;
});

function adjustTextContainerHeight() {
    var windowHeight = window.innerHeight;
    document.getElementById('dynamicTextContainer').style.maxHeight = windowHeight + 'px';
}

function insertFicha(row, col) {

    if (row >= 0 && row < matrix_aux.length && col >= 0 && col < matrix_aux[row].length) {
        matrix[row][col] = 5; 

        const targetCell = document.querySelector(`[data-row="${row}"][data-col="${col}"]`);
        if (targetCell) {    targetCell.className = 'ficha';  }
    } else {
        console.error('Posición fuera de los límites de la matriz.');
    }
}

function MazeTable() {

    mazeContainer.innerHTML = '';
    
    const table = document.createElement('mazeTable');

    matrix_aux.forEach((row, rowIndex) => {
        const tr = document.createElement('tr');
        row.forEach((cellValue, colIndex) => {
            const td = document.createElement('td');
            td.setAttribute('data-row', rowIndex); 
            td.setAttribute('data-col', colIndex);

            if (cellValue === 1) {
                td.className = 'wall';
            } else if (cellValue === 2) {
                td.className = 'trap';
            } else if (cellValue === 3) {
                td.className = 'start';
            } else if (cellValue === 4) {
                td.className = 'end';
            } else if (cellValue === 5) {
                td.className = 'ficha';
            } else if (cellValue === 0) {
                td.className = 'empty';
            }

            tr.appendChild(td);
        });
        table.appendChild(tr);
    });

    mazeContainer.appendChild(table);
}


function initMazeTable() {

    mazeContainer.innerHTML = '';
    
    const table = document.createElement('mazeTable');

    matrix.forEach((row, rowIndex) => {
        const tr = document.createElement('tr');
        row.forEach((cellValue, colIndex) => {
            const td = document.createElement('td');
            td.setAttribute('data-row', rowIndex); 
            td.setAttribute('data-col', colIndex);

            if (cellValue === 1) {
                td.className = 'wall';
            } else if (cellValue === 2) {
                td.className = 'trap';
            } else if (cellValue === 3) {
                td.className = 'start';
            } else if (cellValue === 4) {
                td.className = 'end';
            } else if (cellValue === 5) {
                td.className = 'ficha';
            } else if (cellValue === 0) {
                td.className = 'empty';
            }

            tr.appendChild(td);
        });
        table.appendChild(tr);
    });

    mazeContainer.appendChild(table);
}


