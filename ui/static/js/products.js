"use strict";

const productsTableBody = document.querySelector("tbody");

async function showProducts() {
    try {
        const result = await fetch("http://localhost:4000/admin/panel/read");
        const products = await result.json();


        products.forEach(product => {
            const {code, name, cat, imageUrl, price} = product
            const row = document.createElement("TR");

            const codeEl = document.createElement("TD");
            codeEl.textContent = code;

            const nameEl = document.createElement("TD");
            nameEl.textContent = name;

            const catEl = document.createElement("TD");
            catEl.textContent = cat;

            const imageEl = document.createElement("TD");
            imageEl.textContent = imageUrl;

            const priceEl = document.createElement("TD");
            priceEl.textContent = price

            row.appendChild(codeEl);
            row.appendChild(nameEl);
            row.appendChild(catEl);
            row.appendChild(priceEl);
            row.appendChild(imageEl);

            productsTableBody.appendChild(row)
        });
    } catch (e) {
        console.log(e);
    }
}

showProducts();