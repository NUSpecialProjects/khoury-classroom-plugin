package utils

import "fmt"

func ActionWithDeadlineStub() string {
	var scriptString = `
from datetime import datetime
import sys

def check_date():
    # Define the target date
    target_date = datetime(2024, 12, 20, 0, 0, 0, tzinfo=timezone.utc)
    
    # Get current date and time
    current_date = datetime.now()
    
    # Compare dates and exit with appropriate code
    if current_date > target_date:
        sys.exit(1)
    else:
        sys.exit(0)

if __name__ == "__main__":
    check_date()`

	return fmt.Sprintf(scriptString)
}


func TargetBranchProtectionAction() string {
    var actionString = `name: Check PR Target Branch

on:
  pull_request:
    types: [opened, reopened, edited, synchronize]

jobs:
  check-target:
    runs-on: ubuntu-latest
    steps:
      - name: Check PR destination branch
        run: |
          if [[ "${{ github.event.pull_request.base.ref }}" == "grading" ]]; then
            echo "Error: Pull requests targeting the 'grading' branch are not allowed"
            exit 1
          fi`
          return actionString
}