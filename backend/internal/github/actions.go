package github

import "fmt"

func dateCutScript(date string) string {
	var scriptString = `
from datetime import datetime
import sys

def check_date():
    # Define the target date (November 27, 2024 midnight)
    target_date = datetime(%s)
    
    # Get current date and time
    current_date = datetime.now()
    
    # Compare dates and exit with appropriate code
    if current_date > target_date:
        sys.exit(1)
    else:
        sys.exit(0)

if __name__ == "__main__":
    check_date()`

	return fmt.Sprintf(scriptString, date)
}