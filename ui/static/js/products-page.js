"use mstrict"


async function showAllProducts() {
    try {
        const result = await fetch("http://localhost:4000/admin/panel/read");
        const products = await result.json();

        products.forEach(product => {
            const {code, name, cat, imageUrl, price} = product;

            const card = document.createElement("DIV");
            card.classList.add("shadow-lg", "rounded-2xl", "bg-white", "w-80", "p-2");

            const bodyCard = document.createElement("DIV");
            bodyCard.classList.add("bg-primary", "m-3", "p-4", "rounded-lg");
            bodyCard.setAttribute("value", code);

            const nameEl = document.createElement("P");
            nameEl.textContent = name;
            nameEl.classList.add("text-white", "text-xl", "font-bold");

            const catEl = document.createElement("P");
            catEl.textContent = cat;
            catEl.classList.add("text-gray-50", "text-xs");

            const imageEl = document.createElement("IMG");
            imageEl.setAttribute("src", imageUrl);
            imageEl.classList.add("w-32", "p-4", "h-36", "m-auto")

            const bottomCard = document.createElement("DIV");
            bottomCard.classList.add("flex", "items-center", "justify-between");

            const priceEl = document.createElement("P");
            priceEl.textContent = price + "$";
            priceEl.classList.add("text-white");

            const btn = document.createElement("BUTTON");
            btn.setAttribute("type", "button");
            btn.classList.add("btn", "btn-secondary", "btn-circle", "text-2xl", "font-bold");
            btn.textContent = "+";

            bottomCard.appendChild(priceEl);
            bottomCard.appendChild(btn);

            bodyCard.appendChild(nameEl);
            bodyCard.appendChild(catEl);
            bodyCard.appendChild(bottomCard);

            card.appendChild(imageEl);
            card.appendChild(bodyCard);

            document.querySelector(".flex.flex-wrap.gap-3").appendChild(card);
        });
    } catch (e) {
        console.log(e);
    }
}

showAllProducts();