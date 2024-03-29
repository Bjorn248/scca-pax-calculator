// This file is generated, manual edits will be lost

const PAXIndex = {
  "AM": "1.000",
  "AS": "0.824",
  "BM": "0.978",
  "BS": "0.818",
  "CAM-C": "0.826",
  "CAM-S": "0.848",
  "CAM-T": "0.820",
  "CM": "0.898",
  "CP": "0.856",
  "CS": "0.813",
  "CSM": "0.808",
  "CSP": "0.860",
  "CSX": "0.811",
  "DM": "0.904",
  "DP": "0.865",
  "DS": "0.810",
  "DSP": "0.847",
  "EM": "0.910",
  "EP": "0.858",
  "ES": "0.792",
  "ESP": "0.840",
  "EVX": "0.834",
  "FM": "0.917",
  "FP": "0.877",
  "FS": "0.814",
  "FSAE": "0.980",
  "FSP": "0.831",
  "GS": "0.796",
  "HCR": "0.815",
  "HCS": "0.793",
  "HS": "0.786",
  "KM": "0.937",
  "SM": "0.868",
  "SMF": "0.850",
  "SS": "0.835",
  "SSC": "0.805",
  "SSM": "0.879",
  "SSP": "0.857",
  "SST": "0.836",
  "STH": "0.812",
  "STR": "0.832",
  "STS": "0.816",
  "STU": "0.833",
  "STX": "0.819",
  "XA": "0.844",
  "XB": "0.849",
  "XP": "0.887",
  "XU": "0.869",
};

/**
 * Adds one or more session storage keys to store the answer to a question
 * The index of each array is used to match the key to the value
 * @param {Array} keyArray is an array of key names
 * @param {Array} valueArray is an array of values
 */
function setState(keyArray, valueArray) { // eslint-disable-line no-unused-vars
  for (let i = 0; i < keyArray.length; i++) {
    sessionStorage.setItem(keyArray[i], valueArray[i]);
  }
};

/**
 * This stores the value of an element that triggered an event in sessionStorage
 * @param {event} e is a DOM event
 */
function eventToState(e) {
  setState([e.target.name], [e.target.value]);
};

/**
 * Calculates the required raw time given the
 * competitorRawTime
 * competitorClass
 * yourClass
 * which should be in sessionStorage when this function is called
 */
function calculateRequiredRawTime() {
  // pseudo-code formula
  // competitorPaxTime = rawTime * competitorPax
  // paxTime1 = yourRawTime * yourPax
  // requiredRawTime = paxTime1 / pax2
  rawTime = sessionStorage.getItem("competitorRawTime");
  competitorPax = PAXIndex[sessionStorage.getItem("competitorClass")];
  yourPax = PAXIndex[sessionStorage.getItem("yourClass")];
  competitorPaxParagraph = document.getElementById("competitorPax");
  yourPaxParagraph = document.getElementById("yourPax");
  competitorPaxParagraph.innerHTML = "PAX: " + competitorPax;
  yourPaxParagraph.innerHTML = "PAX: " + yourPax;
  let requiredRawTime = (rawTime * competitorPax) / yourPax;
  // show number with 3 digits after the decimal
  if (!isNaN(requiredRawTime) && requiredRawTime != 0) {
    requiredRawTime = Number(parseFloat(requiredRawTime).toFixed(3));
    const rawTimeToBeat = document.getElementById("rawTimeToBeat");
    rawTimeToBeat.innerHTML = requiredRawTime;
  }
}

/**
 * Populates the select elements with the keys from PAXIndex
 * @param {string} id is the html id of a select element to populate
 */
function populateDropDown(id) {
  const e = document.getElementById(id);
  const classLength = e.options.length;
  for (i = classLength-1; i >= 0; i--) {
    e.remove(i);
  }
  for (const Class of Object.keys(PAXIndex)) {
    const newClass = document.createElement("option");
    newClass.text = Class;
    e.add(newClass);
  }
  if (sessionStorage.getItem(id) in PAXIndex) {
    const providedCompetitorClass = sessionStorage.getItem(id);
    e.value = providedCompetitorClass;
  } else {
    sessionStorage.setItem(id, e.value);
  }
};

/**
 * Populates the competitorRawTime text input field if there is a value already
 * in sessionStorage
 */
function populateInput() {
  const e = document.getElementById("competitorRawTime");
  if (sessionStorage.getItem("competitorRawTime")) {
    const competitorRawTime = sessionStorage.getItem("competitorRawTime");
    e.value = competitorRawTime;
  }
}

/**
 * populates the PAX selection dropdowns
 */
function populatePAXDropdowns() {
  populateDropDown("competitorClass");
  populateDropDown("yourClass");
}

/**
 * swaps the class selections
 */
function swapClasses() {
  if (sessionStorage.getItem("competitorClass") && sessionStorage.getItem("yourClass")) {
    const oldCompetitorClass = sessionStorage.getItem("competitorClass");
    sessionStorage.setItem("competitorClass", sessionStorage.getItem("yourClass"));
    sessionStorage.setItem("yourClass", oldCompetitorClass);
    populatePAXDropdowns();
    populateInput();
    calculateRequiredRawTime();
  }
}

const competitorClass = document.getElementById("competitorClass");
const yourClass = document.getElementById("yourClass");
const competitorRawTime = document.getElementById("competitorRawTime");
const submitButton = document.getElementById("submitButton");
const swapButton = document.getElementById("swapButton");

competitorClass.addEventListener("change", eventToState);
yourClass.addEventListener("change", eventToState);
competitorRawTime.addEventListener("change", eventToState);

competitorClass.addEventListener("change", calculateRequiredRawTime);
yourClass.addEventListener("change", calculateRequiredRawTime);
competitorRawTime.addEventListener("change", calculateRequiredRawTime);
submitButton.addEventListener("click", calculateRequiredRawTime);
swapButton.addEventListener("click", swapClasses);

populatePAXDropdowns();
populateInput();
calculateRequiredRawTime();
