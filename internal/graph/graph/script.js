document.addEventListener("DOMContentLoaded", function () {
  const cy = cytoscape({
    container: document.getElementById("cy"),
    autounselectify: true,
    boxSelectionEnabled: false,
    layout: {
      name: "cola",
      infinite: true,
      fit: false,
    },
    style: [
      {
        selector: "node",
        css: {
          "background-color": "#458588",
          "border-color": "#928374",
          "border-width": 2,
          "label": "data(label)",
          "color": "#ebdbb2",
          "text-valign": "center",
          "text-halign": "right",
          "text-margin-x": 8,
          "font-size": "10px",
          "width": "20px",
          "height": "20px",
          "text-shadow": "1px 1px 2px #000000",
        },
      },
      {
        selector: "edge",
        css: {
          "line-color": "#d79921",
          "target-arrow-color": "#fabd2f",
          "target-arrow-shape": "triangle",
          "curve-style": "bezier",
        },
      },
    ],
    elements: { nodes: [], edges: [] },
  });

  // Fetch initial graph data from the server
  fetch("/data")
    .then((response) => response.json())
    .then((data) => {
      cy.add(data.nodes);
      cy.add(data.edges);
      cy.layout({ name: "cola" }).run();
    })
    .catch((err) => {
      console.error("Failed to load graph data:", err);
    });
});
