#!/bin/bash

# Configure Git to use versioned hooks directory

echo "📦 Installing git hooks..."

# Copy pre-commit hook from hooks/ to .git/hooks/
cp hooks/pre-commit .git/hooks/pre-commit
chmod +x .git/hooks/pre-commit

echo "✅ Git hooks installed successfully!"
echo ""
echo "Now, every time you commit:"
echo "  1. 🔨 Binary will be built automatically"
echo "  2. 📄 output.html will be generated"
echo "  3. 📦 Both files will be added to your commit"
echo ""
echo "To disable: rm .git/hooks/pre-commit"

