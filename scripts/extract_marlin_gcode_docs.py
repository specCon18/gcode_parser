import re
import sys
import os

# Function to parse the file for codes
def parse_codes_from_file(filename):
    # Regular expression pattern to match codes: [ ... ] with any values inside
    pattern = r"codes:\s*\[(.*?)\]"

    codes = []

    try:
        # Open the file for reading
        with open(filename, 'r') as file:
            content = file.read()

            # Match the pattern for codes
            match = re.search(pattern, content)
            if match:
                codes_list = match.group(1)
                codes = [code.strip() for code in codes_list.split(',')]
                # Avoid adding 'g000' as a code
                codes = [code for code in codes if code.lower() != 'g000']
    except FileNotFoundError:
        print(f"Error: The file '{filename}' was not found.")
        return None

    return codes

# Function to parse all top-level parameter tags and optional from the parameters section
def parse_tag_and_optional(filename):
    tag_optional_info = []

    try:
        with open(filename, 'r') as file:
            content = file.read()

            # Match all top-level blocks under 'parameters' section, capturing only the top-level 'tag' and 'optional' fields
            blocks = re.findall(r"-\s*tag:\s*([A-Z])\s*.*?optional:\s*(\S+)(?=\s*-|\n)"  , content, re.DOTALL)

            for block in blocks:
                tag = block[0]
                optional_str = block[1].strip().lower()
                optional = optional_str == 'true'
                tag_optional_info.append((tag, optional))

    except FileNotFoundError:
        print(f"Error: The file '{filename}' was not found.")
        return None

    return tag_optional_info

# Function to process all files in a directory
def process_directory(directory_path):
    # List all files in the directory
    for filename in os.listdir(directory_path):
        file_path = os.path.join(directory_path, filename)

        # Check if the file is a regular file (skip directories)
        if os.path.isfile(file_path):
            print(f"Processing file: {filename}")
            
            # Parse codes from the file
            codes = parse_codes_from_file(file_path)
            if codes:
                print(f"Codes: {codes}")
            else:
                print("No matching pattern for codes found.")

            # Parse tag and optional info from the file
            tag_optional_info = parse_tag_and_optional(file_path)
            if tag_optional_info:
                print("Parameters: ", end="")
                # Format the output as requested
                formatted_params = ", ".join([f"({tag},{optional})" for tag, optional in tag_optional_info])
                print(f"[{formatted_params}]")
            else:
                print("No matching tag and optional information found.")
            print()  # Newline for separation between files

# Main function to handle command-line argument
def main():
    if len(sys.argv) != 2:
        print("Usage: python script.py <directory_path>")
        sys.exit(1)

    directory_path = sys.argv[1]

    if not os.path.isdir(directory_path):
        print(f"Error: '{directory_path}' is not a valid directory.")
        sys.exit(1)

    # Process all files in the directory
    process_directory(directory_path)

# Entry point of the script
if __name__ == "__main__":
    main()

