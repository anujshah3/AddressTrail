let addressesData = null;
let dashboardSectionDisplay = false;
let manageAddressSectionDisplay = false;

async function fetchAddresses() {
  console.log(addressesData);
  if (addressesData !== null) {
    return Promise.resolve(addressesData);
  } else {
    try {
      const response = await fetch("http://localhost:8080/api/users/addresses");
      const data = await response.json();
      addressesData = data;
      return addressesData;
    } catch (error) {
      console.error("Error fetching user addresses:", error);
    }
  }
}

function addTimelineSection() {
  const container = document.createElement("div");
  container.id = "timeline-container";

  const timelineSection = document.createElement("section");
  timelineSection.classList.add("container");

  const leftDiv = document.createElement("div");
  leftDiv.classList.add("left", "left-home", "map-timeline-container");

  const innerTimelineSection = document.createElement("section");
  innerTimelineSection.classList.add("timeline-section");
  innerTimelineSection.id = "timeline-section";

  leftDiv.appendChild(innerTimelineSection);

  const rightDiv = document.createElement("div");
  rightDiv.classList.add("right");

  timelineSection.appendChild(leftDiv);
  timelineSection.appendChild(rightDiv);

  container.appendChild(timelineSection);

  document.body.appendChild(container);
}

function removeTimelineSection() {
  const container = document.getElementById("timeline-container");
  if (container) {
    container.remove();
  }
}

function loadDashboard() {
  fetchAddresses()
    .then((data) => {
      const timelineSection = document.getElementById("timeline-section");
      const ul = document.createElement("ul");
      ul.classList.add("timeline");

      let isLeft = true;

      for (let i = data.length - 1; i >= 0; i--) {
        const address = data[i];
        const li = document.createElement("li");
        const startDate = new Date(address.StartDate);
        const endDate = new Date(address.EndDate);

        const startMonth = startDate.toLocaleString("default", {
          month: "short",
        });
        const endMonth = endDate.toLocaleString("default", {
          month: "short",
        });

        const startMonthYear = `${startMonth} ${startDate.getFullYear()}`;
        const endMonthYear = `${endMonth} ${endDate.getFullYear()}`;

        const addressHTML = `
              <div class="${
                isLeft ? "left-position-address" : "right-position-address"
              }">
                <div class="year-wrapper">
                  <span class="year">${startMonthYear} - ${endMonthYear}</span>
                </div>
                <div class="address">
                  <span class="address-street">${address.Street}${
          address.Unit ? ", " + address.Unit : ""
        }</span>
                  <span class="address-city">${address.City}</span>
                  <span class="address-state">${address.State}${
          address.PostalCode ? ", " + address.PostalCode : ""
        }</span>
                </div>
              </div>
            `;
        li.innerHTML = addressHTML;
        li.dataset.addressId = address.AddressID;
        ul.appendChild(li);

        isLeft = !isLeft;
      }

      timelineSection.appendChild(ul);
    })
    .catch((error) => {
      console.error("Error fetching user addresses:", error);
    });
}

function addManageAddressEventListeners() {
  const modal = document.getElementById("modal");
  const openFormBtn = document.getElementById("openFormBtn");
  const closeModalBtn = document.getElementById("closeModalBtn");
  const addressForm = document.getElementById("addressForm");

  openFormBtn.addEventListener("click", () => {
    modal.style.display = "block";
  });

  closeModalBtn.addEventListener("click", () => {
    modal.style.display = "none";
  });

  window.addEventListener("click", (event) => {
    if (event.target === modal) {
      modal.style.display = "none";
    }
  });

  addressForm.addEventListener("submit", (event) => {
    event.preventDefault();

    const payload = {
      Street: addressForm.elements.street.value,
      Unit: addressForm.elements.unit.value,
      City: addressForm.elements.city.value,
      State: addressForm.elements.state.value,
      PostalCode: addressForm.elements.postalCode.value,
      Country: addressForm.elements.country.value,
      StartDate: addressForm.elements.startDate.value,
      EndDate: addressForm.elements.endDate.value,
    };
    console.log(payload);
    addNewAddress(payload);
    modal.style.display = "none";
    addressesData = null;
    removeManageAddressSection();
    addManageAddressSection();
    loadManageAddresses();
  });
}

function removeManageAddressEventListeners() {
  const modal = document.getElementById("modal");
  const openFormBtn = document.getElementById("openFormBtn");
  const closeModalBtn = document.getElementById("closeModalBtn");
  const addressForm = document.getElementById("addressForm");

  openFormBtn.removeEventListener("click", () => {
    modal.style.display = "block";
  });

  closeModalBtn.removeEventListener("click", () => {
    modal.style.display = "none";
  });

  window.removeEventListener("click", (event) => {
    if (event.target === modal) {
      modal.style.display = "none";
    }
  });

  addressForm.removeEventListener("submit", (event) => {
    event.preventDefault();
  });
}

function addManageAddressSection() {
  const wrapperDiv = document.createElement("div");
  wrapperDiv.id = "manage-address-container";

  const container = document.createElement("section");
  container.classList.add("container", "center-view");

  const addressListContainer = document.createElement("div");
  addressListContainer.classList.add("address-list-container");
  addressListContainer.id = "address-list-container";

  const addButton = document.createElement("button");
  addButton.classList.add("add-new-address");
  addButton.id = "openFormBtn";
  addButton.type = "button";
  addButton.innerHTML = `
    <span class="button__label">Add New Address</span>
  `;

  addressListContainer.appendChild(addButton);

  container.appendChild(addressListContainer);

  const modal = document.createElement("div");
  modal.classList.add("modal");
  modal.id = "modal";

  const modalContent = document.createElement("div");
  modalContent.classList.add("modal-content");

  const closeButton = document.createElement("span");
  closeButton.classList.add("close");
  closeButton.id = "closeModalBtn";
  closeButton.innerHTML = "&times;";

  const heading = document.createElement("h2");
  heading.textContent = "Add New Address";

  const addressForm = document.createElement("form");
  addressForm.classList.add("new-address-form");
  addressForm.id = "addressForm";

  const newAddressSection = document.createElement("section");
  newAddressSection.classList.add("new-address-section");

  newAddressSection.innerHTML = `
    <div class="input-container">
      <label class="new-address-label" for="street">Street</label>
      <input class="new-address-input" type="text" id="street" name="street" />
    </div>
    <div class="input-container-row">
      <div class="left-input-container">
        <label class="new-address-label" for="unit">Unit</label>
        <input class="new-address-input" type="text" id="unit" name="unit" />
      </div>
      <div class="right-input-container">
        <label class="new-address-label" for="city">City</label>
        <input class="new-address-input" type="text" id="city" name="city" />
      </div>
    </div>
    <div class="input-container-row">
      <div class="left-input-container">
        <label class="new-address-label" for="state">State</label>
        <input class="new-address-input" type="text" id="state" name="state" />
      </div>
      <div class="right-input-container">
        <label class="new-address-label" for="postalCode">Postal Code</label>
        <input class="new-address-input" type="text" id="postalCode" name="postalCode" />
      </div>
    </div>
    <div class="input-container-row">
      <div class="left-input-container">
        <label class="new-address-label" for="country">Country</label>
        <input class="new-address-input" type="text" id="country" name="country" />
      </div>
      <div class="right-input-container">
        <label class="new-address-label" for="current">Current</label>
        <input type="checkbox" id="current" name="current" onclick="toggleEndDateInput()" />
      </div>
    </div>
    <div class="input-container-row">
      <div class="left-input-container">
        <label class="new-address-label" for="startDate">Start Date</label>
        <input class="new-address-input" type="date" id="startDate" name="startDate" />
      </div>
      <div class="right-input-container" id="endDateContainer">
        <label class="new-address-label" for="endDate">End Date</label>
        <input class="new-address-input" type="date" id="endDate" name="endDate" />
      </div>
    </div>
  `;

  const saveContainer = document.createElement("div");
  saveContainer.classList.add("save-container");

  const saveButton = document.createElement("input");
  saveButton.id = "submit-new-address";
  saveButton.type = "submit";
  saveButton.value = "Save";
  saveButton.onclick = () => true;

  saveContainer.appendChild(saveButton);

  addressForm.appendChild(newAddressSection);
  addressForm.appendChild(saveContainer);

  modalContent.appendChild(closeButton);
  modalContent.appendChild(heading);
  modalContent.appendChild(addressForm);

  modal.appendChild(modalContent);

  document.body.appendChild(modal);

  wrapperDiv.appendChild(container);

  document.body.appendChild(wrapperDiv);
  addManageAddressEventListeners();
}

function removeManageAddressSection() {
  const container = document.getElementById("manage-address-container");
  if (container) {
    removeManageAddressEventListeners();
    container.remove();
  }
}

function toggleEndDateInput() {
  var endDateContainer = document.getElementById("endDateContainer");
  var currentCheckbox = document.getElementById("current");

  if (currentCheckbox.checked) {
    endDateContainer.style.display = "none";
  } else {
    endDateContainer.style.display = "block";
  }
}

function setDefaultStartDate() {
  var startDateInput = document.getElementById("startDate");
  var endDateInput = document.getElementById("endDate");

  var today = new Date();
  var year = today.getFullYear();
  var month = String(today.getMonth() + 1).padStart(2, "0");
  var day = String(today.getDate()).padStart(2, "0");

  var todayFormatted = year + "-" + month + "-" + day;
  startDateInput.value = todayFormatted;
  endDateInput.value = todayFormatted;
}

function loadManageAddresses() {
  fetchAddresses()
    .then((data) => {
      const addressListContainer = document.getElementById(
        "address-list-container"
      );

      data.reverse().forEach((address) => {
        const addressHolder = document.createElement("div");
        addressHolder.classList.add("address-holder");

        const addressLine = document.createElement("button");
        addressLine.classList.add("address-line");
        addressLine.dataset.addressId = address.AddressID;

        const streetAddressDisplay = document.createElement("div");
        streetAddressDisplay.classList.add("street-address-display");
        streetAddressDisplay.innerText = address.Street;

        const cityStateAddressDisplay = document.createElement("div");
        cityStateAddressDisplay.classList.add("city-state-address-display");

        const startDate = new Date(address.StartDate);
        const endDate = new Date(address.EndDate);

        const startMonthYear = `${startDate.toLocaleString("default", {
          month: "short",
        })} ${startDate.getFullYear()}`;
        const endMonthYear = `${endDate.toLocaleString("default", {
          month: "short",
        })} ${endDate.getFullYear()}`;

        cityStateAddressDisplay.innerText = `${address.City}, ${address.State} ${startMonthYear} - ${endMonthYear}`;

        const editIcon = document.createElement("div");
        editIcon.classList.add(
          "apply-flow-profile-item-tile__edit-icon",
          "background-color-1"
        );

        addressLine.appendChild(streetAddressDisplay);
        addressLine.appendChild(cityStateAddressDisplay);
        addressLine.appendChild(editIcon);
        addressHolder.appendChild(addressLine);
        addressListContainer.appendChild(addressHolder);
      });
    })
    .catch((error) => {
      console.error("Error fetching user addresses:", error);
    });
}

async function addNewAddress(addressData) {
  try {
    const response = await fetch("http://localhost:8080/api/users/addresses", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(addressData),
    });

    if (response.ok) {
      const data = await response.json();
      console.log(data);
      return data;
    } else {
      throw new Error("Request failed with status: " + response.status);
    }
  } catch (error) {
    console.error(error);
    throw error;
  }
}

document.addEventListener("DOMContentLoaded", function () {
  addTimelineSection();
  loadDashboard();
  dashboardSectionDisplay = true;

  document
    .getElementById("dashboard-link")
    .addEventListener("click", function (event) {
      event.preventDefault();
      if (manageAddressSectionDisplay) {
        removeManageAddressSection();
        manageAddressSectionDisplay = false;
      }
      if (!dashboardSectionDisplay) {
        addTimelineSection();
        loadDashboard();
        dashboardSectionDisplay = true;
      }
    });
  document
    .getElementById("manage-addresses-link")
    .addEventListener("click", function (event) {
      event.preventDefault();
      if (dashboardSectionDisplay) {
        removeTimelineSection();
        dashboardSectionDisplay = false;
      }
      if (!manageAddressSectionDisplay) {
        addManageAddressSection();
        loadManageAddresses();
        manageAddressSectionDisplay = true;
      }
    });
});
