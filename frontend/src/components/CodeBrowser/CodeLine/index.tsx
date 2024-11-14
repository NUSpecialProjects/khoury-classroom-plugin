import { useContext, useEffect, useRef, useState } from "react";

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

  const points = useRef<HTMLInputElement>(null);

  useEffect(() => {
    setEditing(false);
  }, [path]);

  const adjustPoints = (x: number) => {
    if (!points.current) return;
    const pts = points.current.value;
    points.current.value = (parseInt(pts, 10) + x).toString();
  };

  const saveComment = (e: React.FormEvent) => {
    e.preventDefault();

    if (!selectedClassroom || !assignmentID || !studentWorkID) return;
    const form = e.target as HTMLFormElement;
    const data = new FormData(form);
    const comment: IGradingComment = {
      path,
      line: Number(data.get("line")),
      body: String(data.get("comment")).trim(),
      points: Number(data.get("points")),
    };
    if (comment.points == 0 && comment.body == "") return;
    if (comment.body == "")
      comment.body = "No comment left for this point adjustment.";
    addComment(comment);
    setEditing(false);
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
              const pointSign =
                comment.points > 0
                  ? "positive"
                  : comment.points < 0
                    ? "negative"
                    : "neutral";
              return (
                <div
                  className="CodeLine__commentWrapper CodeLine__comment"
                  key={Number(i)}
                >
                  <div
                    className={`CodeLine__commentPoints CodeLine__commentPoints--${pointSign}`}
                  >
                    {comment.points == 0
                      ? "Comment"
                      : comment.points > 0
                        ? `+${comment.points}`
                        : comment.points}
                  </div>
                  {comment.body}
                </div>
              );
            })}

          {/************ Display form to create new comment *************/}
          {editing && (
            <div className="CodeLine__commentWrapper">
              <form className="CodeLine__newCommentForm" onSubmit={saveComment}>
                <input readOnly hidden type="number" name="line" value={line} />

                <div className="CodeLine__newCommentPoints">
                  <label htmlFor="points">Point Adjustment</label>
                  <input
                    ref={points}
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
