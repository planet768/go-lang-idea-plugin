package main
var c = "asdfasdf", "adfad\"adfadsf", "\\\n\a\b\f\n\r\t\v", `abc`, "\n", "", "Hello, world!\n", "日本語", "\u65e5本\U00008a9e", "\xff\u00FF", `\n
\n`
func a() { j = i }
/**-----
Go file
  PackageDeclaration(main)
    PsiElement(KEYWORD_PACKAGE)('package')
    PsiWhiteSpace(' ')
    PsiElement(IDENTIFIER)('main')
  PsiWhiteSpace('\n')
  VarDeclarationsImpl
    PsiElement(KEYWORD_VAR)('var')
    PsiWhiteSpace(' ')
    VarDeclarationImpl
      LiteralIdentifierImpl
        PsiElement(IDENTIFIER)('c')
      PsiWhiteSpace(' ')
      PsiElement(=)('=')
      PsiWhiteSpace(' ')
      LiteralExpressionImpl
        LiteralStringImpl
          PsiElement(LITERAL_STRING)('"asdfasdf"')
      PsiElement(,)(',')
      PsiWhiteSpace(' ')
      LiteralExpressionImpl
        LiteralStringImpl
          PsiElement(LITERAL_STRING)('"adfad\"adfadsf"')
      PsiElement(,)(',')
      PsiWhiteSpace(' ')
      LiteralExpressionImpl
        LiteralStringImpl
          PsiElement(LITERAL_STRING)('"\\\n\a\b\f\n\r\t\v"')
      PsiElement(,)(',')
      PsiWhiteSpace(' ')
      LiteralExpressionImpl
        LiteralStringImpl
          PsiElement(LITERAL_STRING)('`abc`')
      PsiElement(,)(',')
      PsiWhiteSpace(' ')
      LiteralExpressionImpl
        LiteralStringImpl
          PsiElement(LITERAL_STRING)('"\n"')
      PsiElement(,)(',')
      PsiWhiteSpace(' ')
      LiteralExpressionImpl
        LiteralStringImpl
          PsiElement(LITERAL_STRING)('""')
      PsiElement(,)(',')
      PsiWhiteSpace(' ')
      LiteralExpressionImpl
        LiteralStringImpl
          PsiElement(LITERAL_STRING)('"Hello, world!\n"')
      PsiElement(,)(',')
      PsiWhiteSpace(' ')
      LiteralExpressionImpl
        LiteralStringImpl
          PsiElement(LITERAL_STRING)('"日本語"')
      PsiElement(,)(',')
      PsiWhiteSpace(' ')
      LiteralExpressionImpl
        LiteralStringImpl
          PsiElement(LITERAL_STRING)('"\u65e5本\U00008a9e"')
      PsiElement(,)(',')
      PsiWhiteSpace(' ')
      LiteralExpressionImpl
        LiteralStringImpl
          PsiElement(LITERAL_STRING)('"\xff\u00FF"')
      PsiElement(,)(',')
      PsiWhiteSpace(' ')
      LiteralExpressionImpl
        LiteralStringImpl
          PsiElement(LITERAL_STRING)('`\n\n\n`')
  PsiWhiteSpace('\n')
  FunctionDeclaration(a)
    PsiElement(KEYWORD_FUNC)('func')
    PsiWhiteSpace(' ')
    LiteralIdentifierImpl
      PsiElement(IDENTIFIER)('a')
    PsiElement(()('(')
    PsiElement())(')')
    PsiWhiteSpace(' ')
    BlockStmtImpl
      PsiElement({)('{')
      PsiWhiteSpace(' ')
      AssignStmtImpl
        ExpressionListImpl
          LiteralExpressionImpl
            LiteralIdentifierImpl
              PsiElement(IDENTIFIER)('j')
        PsiWhiteSpace(' ')
        PsiElement(=)('=')
        PsiWhiteSpace(' ')
        ExpressionListImpl
          LiteralExpressionImpl
            LiteralIdentifierImpl
              PsiElement(IDENTIFIER)('i')
      PsiWhiteSpace(' ')
      PsiElement(})('}')
