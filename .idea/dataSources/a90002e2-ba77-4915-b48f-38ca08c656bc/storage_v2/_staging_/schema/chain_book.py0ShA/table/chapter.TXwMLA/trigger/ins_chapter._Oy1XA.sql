create trigger ins_chapter
  before INSERT
  on chapter
  for each row
  begin
    declare content int;
    SET content = (SELECT chapter.sort
                   FROM chain_book.chapter
                   WHERE chapter.book_id = NEW.book_id
                     and chapter.sort = NEW.sort);
    if content > 0
    then
      delete
      from chapter
      where book_id = NEW.book_id
        and sort = NEW.sort;
    end if;
  end;