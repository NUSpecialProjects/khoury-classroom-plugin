import { useContext, useState } from "react";

import { GraderContext } from "@/contexts/grader";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";

import "./styles.css";
import Button from "@/components/Button";

interface ICodeLine extends React.HTMLProps<HTMLDivElement> {
  path: string;
  line: number;
  isDiff: boolean;
  code: string;
}

const CodeLine: React.FC<ICodeLine> = ({ path, line, isDiff, code }) => {
  const { selectedClassroom } = useContext(SelectedClassroomContext);
  const { assignmentID, studentWorkID, comments, addComment } =
    useContext(GraderContext);
  const [editing, setEditing] = useState(false);

  const adjustPoints = (x: number) => {
    const input = document.getElementById("points") as HTMLInputElement;
    input.value = (parseInt(input.value, 10) + x).toString();
  };

  const saveComment = (e: React.FormEvent) => {
    e.preventDefault();

    if (!selectedClassroom || !assignmentID || !studentWorkID) return;
    const form = e.target as HTMLFormElement;
    const data = new FormData(form);
    const comment: IGradingComment = {
      path,
      line: Number(data.get("line")),
      body: String(data.get("comment")),
      points: Number(data.get("points")),
    };
    addComment(comment);
    form.reset();
  };

  return (
    <>
      <div className={`CodeLine${isDiff ? " CodeLine--diff" : ""}`}>
        <div className="CodeLine__number">
          {line}
          {isDiff && (
            <div
              className="CodeLine__newCommentButton"
              onClick={() => {
                setEditing(!editing);
              }}
            >
              +
            </div>
          )}
        </div>
        <div
          className="CodeLine__content"
          dangerouslySetInnerHTML={{ __html: code }}
        ></div>
      </div>
      {(editing || comments[path]?.[line]) && (
        <div className="CodeLine__comments">
          {/************ Display any existing comments *************/}
          {comments[path]?.[line] &&
            Object.entries(comments[path][line]).map(([i, comment]) => {
              return (
                <div className="CodeLine__comment" key={Number(i)}>
                  {comment.body}
                </div>
              );
            })}

          {/************ Display form to create new comment *************/}
          {editing && (
            <div style={{ display: "flex", justifyContent: "space-between" }}>
              <form className="CodeLine__newCommentForm" onSubmit={saveComment}>
                <input readOnly hidden type="number" name="line" value={line} />

                <div className="CodeLine__newCommentPoints">
                  <label htmlFor="points">Point Adjustment</label>
                  <input
                    id="points"
                    type="number"
                    name="points"
                    defaultValue={0}
                  />
                  <div className="CodeLine__newCommentPoints__spinners">
                    <div
                      tabIndex={0}
                      onClick={() => {
                        adjustPoints(1);
                      }}
                    >
                      +
                    </div>
                    <div
                      tabIndex={0}
                      onClick={() => {
                        adjustPoints(-1);
                      }}
                    >
                      -
                    </div>
                  </div>
                </div>

                <textarea name="comment" placeholder="Leave a comment" />
                <div className="CodeLine__newCommentButtons">
                  <Button
                    className="CodeLine__newCommentCancel"
                    onClick={() => {
                      setEditing(false);
                    }}
                  >
                    Cancel
                  </Button>
                  <Button type="submit">Save</Button>
                </div>
              </form>
            </div>
          )}
        </div>
      )}
    </>
  );
};

export default CodeLine;
