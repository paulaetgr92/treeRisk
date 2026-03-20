// Inicializa mapa (São Paulo)
const map = L.map('map').setView([-23.55, -46.63], 13);

// Camada do mapa
L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '&copy; OpenStreetMap'
}).addTo(map);

let clickEnabled = false;
let trees = [];

// Ativar clique no mapa
function enableMapClick() {
    clickEnabled = true;
    document.getElementById("status").innerText = "Clique no mapa para cadastrar a árvore...";
}

// Evento de clique no mapa
map.on('click', function(e) {

    if (!clickEnabled) return;

    const species = document.getElementById("species").value;
    const height = document.getElementById("height").value;

    if (!species || !height) {
        alert("Preencha todos os campos!");
        return;
    }

    const lat = e.latlng.lat;
    const lng = e.latlng.lng;

    // Adiciona marcador
    const marker = L.marker([lat, lng]).addTo(map);

    marker.bindPopup(`
    <b>${species}</b><br>
    Altura: ${height}m
  `);

    // Salva árvore
    const tree = { species, height, lat, lng };
    trees.push(tree);

    updateTreeList();

    document.getElementById("status").innerText = "Árvore cadastrada com sucesso!";
    clickEnabled = false;

    // Limpa inputs
    document.getElementById("species").value = "";
    document.getElementById("height").value = "";
});

// Atualiza lista na tela
function updateTreeList() {
    const list = document.getElementById("tree-list");

    list.innerHTML = "<h3>Árvores cadastradas:</h3>";

    trees.forEach((tree, index) => {
        list.innerHTML += `
      <div>
        🌳 ${tree.species} - ${tree.height}m 
        <br>
        📍 (${tree.lat.toFixed(4)}, ${tree.lng.toFixed(4)})
      </div>
      <hr>
    `;
    });
}

// Expandir mapa
function toggleMap() {
    const mapDiv = document.getElementById("map");

    if (mapDiv.style.height === "90vh") {
        mapDiv.style.height = "400px";
    } else {
        mapDiv.style.height = "90vh";
    }

    // Corrige renderização do Leaflet
    setTimeout(() => {
        map.invalidateSize();
    }, 200);
}