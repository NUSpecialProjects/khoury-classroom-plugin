import { FaChevronLeft, FaChevronRight } from "react-icons/fa";
import { Light as CodeViewer } from "react-syntax-highlighter";
import { hybrid } from "react-syntax-highlighter/dist/esm/styles/hljs";

import Button from "@/components/Button";

import "./styles.css";

const Grader: React.FC = () => {
  const codeText = `# Python Program to find the L.C.M. of two input number

def compute_lcm(x, y):

    # choose the greater number
    if x > y:
        greater = x
    else:
        greater = y

    while(True):
        if((greater % x == 0) and (greater % y == 0)):
            lcm = greater
            break
        greater += 1

    return lcm

num1 = 54
num2 = 24

print("The L.C.M. is", compute_lcm(num1, num2))
`;

  return (
    <div className="Grader">
      <div className="Grader__head">
        <div className="Grader__title">
          <FaChevronLeft />
          <div>
            <h2>Assignment 3</h2>
            <span>Jane Doe</span>
          </div>
        </div>
        <div className="Grader__nav">
          <span>Submission 2/74</span>
          <div>
            <Button>
              <FaChevronLeft />
              Previous
            </Button>
            <Button>
              Next
              <FaChevronRight />
            </Button>
          </div>
        </div>
      </div>
      <div className="Grader__body">
        <div className="Grader__files">
          <div className="Grader__file">file.txt</div>
          <div className="Grader__file">file.txt</div>
          <div className="Grader__file">file.txt</div>
        </div>
        <CodeViewer
          className="Grader__code"
          showLineNumbers
          lineNumberStyle={{ color: "#999", margin: "0 5px" }}
          language="python"
          style={hybrid}
        >
          {codeText}
        </CodeViewer>
      </div>
    </div>
  );
};

export default Grader;
