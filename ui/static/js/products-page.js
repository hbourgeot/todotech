"use strict";

let productsObj;
let cart = [];
let cost = 0;

const productsBtns = document.querySelectorAll(".btn.btn-circle");
const totalCost = document.querySelector("#total-cost");

async function storeProducts() {
  try {
    const result = await fetch("https://todotech.ml/admin/read");
    const products = await result.json();
    productsObj = JSON.parse(JSON.stringify(products));
  } catch (e) {
    console.log(e);
  }
}

productsBtns.forEach((btn) => {
  btn.addEventListener("click", addtoCart);
});

function addtoCart(e) {
  cart.push(e.target.getAttribute("product-code"));
  renderCart();
}
function renderCart() {
  document.querySelector("#shopping-cart").textContent = "";
  const cartwithoutDupl = [...new Set(cart)];
  cartwithoutDupl.forEach((product) => {
    const myProduct = productsObj.filter((item) => {
      return item.code === parseInt(product);
    });

    const numProducts = cart.reduce((total, productId) => {
      return productId === product ? (total += 1) : total;
    }, 0);

    const productRow = document.createElement("DIV");
    productRow.classList.add(
      "flex",
      "items-center",
      "hover:bg-gray-100",
      "py-5"
    );

    const productDiv = document.createElement("DIV");
    productDiv.classList.add("flex", "w-2/5");

    const prodImgDiv = document.createElement("DIV");
    prodImgDiv.classList.add("w-20");

    const productImg = document.createElement("IMG");
    productImg.setAttribute("src", myProduct[0].imageUrl);
    productImg.classList.add("h-24");

    prodImgDiv.appendChild(productImg);

    const prodDetDiv = document.createElement("DIV");
    prodDetDiv.classList.add(
      "flex",
      "flex-col",
      "justify-between",
      "ml-4",
      "flex-grow"
    );

    const productName = document.createElement("SPAN");
    productName.classList.add("font-bold", "text-sm");
    productName.textContent = myProduct[0].name;

    const productCat = document.createElement("SPAN");
    productCat.classList.add("text-red-500", "text-xs");
    productCat.textContent = myProduct[0].cat;

    const prodBtnDel = document.createElement("BUTTON");
    prodBtnDel.classList.add(
      "font-semibold",
      "hover:text-red-500",
      "text-gray-500",
      "text-xs"
    );
    prodBtnDel.textContent = "Remove";
    prodBtnDel.dataset.item = product;
    prodBtnDel.addEventListener("click", deleteProductCart);

    prodDetDiv.appendChild(productName);
    prodDetDiv.appendChild(productCat);
    prodDetDiv.appendChild(prodBtnDel);

    productDiv.appendChild(prodImgDiv);
    productDiv.appendChild(prodDetDiv);

    const productQuant = document.createElement("INPUT");
    productQuant.setAttribute("readonly", "");
    productQuant.classList.add(
      "w-1/5",
      "mx-2",
      "border",
      "text-center",
      "font-semibold",
      "text-sm",
      "input",
      "input-ghost"
    );
    productQuant.value = numProducts;

    const productPrice = document.createElement("INPUT");
    productPrice.classList.add(
      "w-1/5",
      "mx-2",
      "border",
      "text-center",
      "font-semibold",
      "input",
      "input-ghost"
    );
    productPrice.setAttribute("readonly", "");
    productPrice.value = myProduct[0].price + " USD$";

    const totalProductPrice = document.createElement("INPUT");
    totalProductPrice.classList.add(
      "w-1/5",
      "mx-2",
      "border",
      "text-center",
      "font-semibold",
      "text-sm",
      "input",
      "input-ghost"
    );
    totalProductPrice.setAttribute("readonly", "");
    totalProductPrice.value =
      parseInt(productPrice.value) * parseInt(numProducts) + " USD$";

    productRow.appendChild(productDiv);
    productRow.appendChild(productQuant);
    productRow.appendChild(productPrice);
    productRow.appendChild(totalProductPrice);

    document.querySelector("#shopping-cart").appendChild(productRow);

    cost += parseInt(productPrice.value);
  });
  totalCost.textContent = calcTotal();
}

function deleteProductCart(e) {
  const code = e.target.dataset.item;

  cart = cart.filter((cartCode) => {
    return cartCode !== code;
  });
  renderCart();
}

function calcTotal() {
  return cart
    .reduce((total, productCart) => {
      const myProduct = productsObj.filter((product) => {
        return product.code === parseInt(productCart);
      });
      return total + myProduct[0].price;
    }, 0)
    .toFixed(2);
}

document.getElementById("checkout").addEventListener("click", () => {
  const name = document.getElementById("order-name").value;
  const userName = document.getElementById("order-username").value;
  const country = document.getElementById("order-country").value;

  const sendCart = [...cart];

  const order = {
    name: name,
    userName: userName,
    country: country,
    cart: sendCart,
  };

  fetch("https://todotech.ml/neworder", {
    method: "POST",
    body: JSON.stringify(order),
    headers: {
      "Content-type": "application/json",
    },
  }).then((window.location = "/products"));
});

storeProducts();
