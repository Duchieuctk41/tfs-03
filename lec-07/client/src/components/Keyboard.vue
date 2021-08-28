<template>
  <div class="keyboard">
    <button class="key --gray-dark" @click="clear">C</button>
    <button class="key --gray-dark" @click="deleteIn">DE</button>
    <button class="key --gray-dark" @click="addOpe('%')">%</button>
    <button class="key --gray-dark" @click="addOpe('/')">/</button>
    <button class="key" @click="addNum('7')">7</button>
    <button class="key" @click="addNum('8')">8</button>
    <button class="key" @click="addNum('9')">9</button>
    <button class="key --yellow" @click="addOpe('*')">X</button>
    <button class="key" @click="addNum('4')">4</button>
    <button class="key" @click="addNum('5')">5</button>
    <button class="key" @click="addNum('6')">6</button>
    <button class="key --yellow" @click="addOpe('-')">-</button>
    <button class="key" @click="addNum('1')">1</button>
    <button class="key" @click="addNum('2')">2</button>
    <button class="key" @click="addNum('3')">3</button>
    <button class="key --yellow" @click="addOpe('+')">+</button>
    <button class="key col-2" @click="addNum('0')">0</button>
    <button class="key" @click="addComma">.</button>
    <button class="key --yellow" @click.prevent="submit">=</button>
  </div>
</template>

<script>
import axios from "axios";

let dataPush = {
  operators: [],
  numbers: [],
};

export default {
  name: "Keyboard",
  data() {
    return {
      isClear: false,
      isComma: true,
      isTrue: false, // check input nhập trước đó là số hay toán tử, để thêm " " cho đẹp 
      input: "0",
      output: "Result",
    };
  },
  methods: {
    addOpe: function (addOpe) {
      if (this.isClear) this.clear();
      if (this.isTrue) this.input += " " + addOpe + " ";
      this.isTrue = false;
      this.isComma = true;
      this.$emit("update", {input: this.input, output: this.output});
    },

    addNum: function (number) {
      if (this.isClear) this.clear();
      if (this.input == "0") this.input = "";
      this.input += number;
      this.isTrue = true;
      this.$emit("update", {input: this.input, output: this.output});
    },

    clear: function () {
      dataPush = {
        operators: [],
        numbers: [],
      };
      (this.input = "0"), (this.output = "Result"), (this.isTrue = false);
      this.isClear = false;
      this.isComma = true;
      this.$emit("update", {input: this.input, output: this.output});
    },

    addComma: function () {
      if (this.isClear) clear();
      if (this.isComma) {
        this.input += ".";
        this.isComma = false;
      }
      this.$emit("update", {input: this.input, output: this.output});
    },
    submit: async function () {
      this.isClear = true;
      this.isTrue = false;
      let substrings = this.input.split(" ");
      for (let i = 0; i < substrings.length; i++) {
        i % 2 === 1
          ? dataPush.operators.push(substrings[i])
          : dataPush.numbers.push(Number(substrings[i]));
      }
      const { data } = await axios.post(
        "http://localhost:3000/api/calc",
        dataPush
      );
      this.$emit("update", {input: this.input, output: data.msg});
    },

    deleteIn: function () {
      if (this.input.length === 1) this.clear();
      if (this.input !== "0") {
        let lastElement = this.input[this.input.length - 1];
        if (lastElement === " ") {
          this.input = this.input.substr(0, this.input.length - 2);
          this.isTrue = true;
        }
        this.input = this.input.substr(0, this.input.length - 1);
      }
      this.$emit("update", {input: this.input, output: this.output});
    },
  },
};
</script>

<style scoped>
.keyboard {
  display: grid;
  grid-template-columns: repeat(4, 90px);
  gap: 10px;
  margin: 0 auto;
  padding: 0 5px;
}

.key {
  height: 80px;
  width: 80px;
  font-size: 2rem;
  font-weight: bold;
  cursor: pointer;
  background-color: #4d4d4d;
  color: #fff;
  border: none;
  box-shadow: 0 1px 4px rgb(0 0 0 / 16%);
  border-radius: 8px;
}

.key:hover {
  filter: brightness(1.1);
}

.key:focus {
  animation: pressBtn 0.3s ease-in-out;
}

.key.col-2 {
  grid-column: 1/3;
  width: 180px;
}

.key.--yellow {
  background-color: #ff9323;
}

.key.--gray-dark {
  background-color: #afafaf;
  color: #2e2e2e;
}

@keyframes pressBtn {
  50% {
    position: relative;
    top: 1px;
  }
  100% {
    position: inherit;
    top: 0;
  }
}
</style>