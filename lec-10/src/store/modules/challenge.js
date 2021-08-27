const state = {
  title: "Order summary component",
  description:
    "A perfect project for newbies who are starting to build confidence with layouts!",
  lang: ["HTML", "CSS"],
  price: "free",
  level: {
    number: 1,
    name: "Newbie",
  },
  asset: [
    {
      name: "A perfect project for newbies who are starting to build confidence with layouts!",
      status: false,
      id: "asset_1",
    },
    {
      name: "JPEG design files for mobile & desktop layouts",
      status: true,
      id: "asset_2",
    },
    { name: "Style guide for fonts, colors, etc", status: true, id: "asset_3" },
    { name: "Optimized image assets", status: true, id: "asset_4" },
    {
      name: "README file to help you get started",
      status: true,
      id: "asset_5",
    },
    { name: "HTML file with pre-written content", status: true, id: "asset_6" },
  ],
  image: [
    { name: "desktop design", img: "desktop.jpg"},
    { name: "mobile design", img: "mobile.jpg" },
  ],
  brief: [
    {
      content:
        "Your challenge is to build out this order summary card component and get it looking as close to the design as possible.",
      type: "p",
    },
    {
      content:
        "You can use any tools you like to help you complete the challenge. So if you've got something you'd like to practice, feel free to give it a go.",
      type: "p",
    },
    { content: "Your users should be able to:", type: "p" },
    { content: "See hover states for interactive elements", type: "ul" },
    {
      content:
        "Download the project and go through the README.md file. This will provide further details about the project and help you get set up.",
      type: "p",
    },
    {
      content:
        "Want some support on the challenge? Join our Slack community and ask questions in the help channel.",
      type: "p",
    },
  ],
  guide: [
    { content: "Download the starter code" },
    { content: "Set up the project with version control (e.g. Git)" },
    { content: "Read the README.md file and have a look around the project" },
    { content: "Get colors, fonts etc from the style-guide.md file" },
    { content: "Set up your project/file architecture however you want" },
    { content: "Start coding!" },
  ],
  idea: [
    {
      content:
        "Write your styles using a pre-processor, such as Sass, Less or Stylus",
    },
    {
      content:
        "Train your eye for detail by getting your solution as close to the design as you can",
    },
    {
      content:
        "Try estimating the time it will take for you to build the project. Then see if the time taken matches up to your estimate. Project estimations are a skill that is often overlooked, but is important for professional developers",
    },
  ],
  faq: [
    {
      sum: "Can I use libraries/frameworks on these projects?",
      desc: "Yes! Our challenges provide professional designs but there are no rules on what tools to use. So feel free to use anything you like to build your projects.",
    },
    {
      sum: "How can I get help if I'm stuck on a challenge?",
      desc: "The best (and quickest) way to get help on a challenge is in our Slack community. There are thousands of other developers in there, so it's a great place to ask questions. We even have a dedicated help channel! If you haven't joined yet, you can get an invite to our Slack community here.",
    },
    {
      sum: "Can I use these projects in my portfolio?",
      desc: "Definitely! Please do feel free to use whatever you build in your portfolio. Helping developers add professional-looking projects to their portfolio was one of the reasons we created this platform!",
    },
    {
      sum: "Is there an official solution I can take a look at?",
      desc: "We don't provide official solutions for the challenges. This is because there is no single perfect way to complete a challenge. Instead, you're encouraged to review other people's code in the community. You can learn so much by seeing how other people have approached the same challenges and giving them feedback.",
    },
    {
      sum: "Do I get a code review when I post my solution?",
      desc: "Frontend Mentor is a collaborative learning community where everyone can give feedback to each other. If you'd like to receive feedback from the community, please be sure to post a question when you submit your solution. The more specific you can be, the better. Being clear with your questions means you're much more likely to receive valuable feedback from others.",
    },
    {
      sum: "How do I submit my solution?",
      desc: "We'd recommend reading our complete guide to submitting solutions. If you get stuck and need help, please feel free to ask questions in our Slack community, and we'll help you submit your project.",
    },
    {
      sum: "Can I use these challenges within my own free or commercial content/tutorials/projects?",
      desc: "Please visit our license page to learn more about how our challenges can be used within your own content. If you're unsure about anything, please feel free to contact us at hi@frontendmentor.io and we'll be more than happy to answer your questions.",
    },
  ],
};

const getters = {
  allContent: (state) => state,
};

export default {
  state,
  getters,
};
