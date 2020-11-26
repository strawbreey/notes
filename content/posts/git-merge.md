---
title: "Git Merge"
date: 2020-11-25T17:20:21+08:00
draft: false
---



在`git merge`命令中，有以下三种使用参数：

```
git merge [-n] [--stat] [--no-commit] [--squash] [--[no-]edit] [-s <strategy>] [-X <strategy-option>] [-S[<keyid>]] [--[no-]rerere-autoupdate] [-m <msg>] [<commit>...]
git merge <msg> HEAD <commit>...
git merge --abort
```

### git-merge简介

git-merge命令是用于从指定的commit(s)合并到当前分支的操作。

> 注：这里的指定commit(s)是指从这些历史commit节点开始，一直到当前分开的时候。

git-merge命令有以下两种用途：

- 用于git-pull中，来整合另一代码仓库中的变化（即：git pull = git fetch + git merge）

- 用于从一个分支到另一个分支的合并


1. git merge <msg> HEAD <commit>...命令

该命令的存在是由于历史原因，在新版本中不应该使用它，应该使用git merge -m <msg> <commit>....进行替代

2. git merge --abort命令
该命令仅仅在合并后导致冲突时才使用。git merge --abort将会抛弃合并过程并且尝试重建合并前的状态。但是，当合并开始时如果存在未commit的文件，git merge --abort在某些情况下将无法重现合并前的状态。（特别是这些未commit的文件在合并的过程中将会被修改时）

警告：运行git-merge时含有大量的未commit文件很容易让你陷入困境，这将使你在冲突中难以回退。因此非常不鼓励在使用git-merge时存在未commit的文件，建议使用git-stash命令将这些未commit文件暂存起来，并在解决冲突以后使用git stash pop把这些未commit文件还原出来。

### git 参数

`git merge` 命令中使用的参数

1. --commit和--no-commit

--commit参数使得合并后产生一个合并结果的commit节点。该参数可以覆盖--no-commit。
--no-commit参数使得合并后，为了防止合并失败并不自动提交，能够给使用者一个机会在提交前审视和修改合并结果。

2. --edit|-e 以及 --no-edit

--edit和-e用于在成功合并、提交前调用编辑器来进一步编辑自动生成的合并信息。因此使用者能够进一步解释和判断合并的结果。
--no-edit参数能够用于接受自动合并的信息（通常情况下并不鼓励这样做）。

如果你在合并时已经给定了-m参数（下文介绍），使用 --edit（或-e）依然是有用的，这将在编辑器中进一步编辑-m所含的内容。


3. --ff命令

--ff是指fast-forward命令。当使用fast-forward模式进行合并时，将不会创造一个新的commit节点。默认情况下，git-merge采用fast-forward模式。
关于fast-forward模式的详细解释，请看我的另一篇文章：一个成功的Git分支模型的“关于fast forward”一节。

4. --no-ff命令

即使可以使用fast-forward模式，也要创建一个新的合并节点。这是当git merge在合并一个tag时的默认行为。

5. --ff-only命令

除非当前HEAD节点已经up-to-date（更新指向到最新节点）或者能够使用fast-forward模式进行合并，否则的话将拒绝合并，并返回一个失败状态。

6. --log[=<n>]和 --no-log
--log[=<n>]将在合并提交时，除了含有分支名以外，还将含有最多n个被合并commit节点的日志信息。
--no-log并不会列出该信息。

7. --stat, -n, --no-stat命令
--stat参数将会在合并结果的末端显示文件差异的状态。文件差异的状态也可以在git配置文件中的merge.stat配置。
相反，-n, --no-stat参数将不会显示该信息。

8. --squash 和--no-squash
--squash 当一个合并发生时，从当前分支和对方分支的共同祖先节点之后的对方分支节点，一直到对方分支的顶部节点将会压缩在一起，使用者可以经过审视后进行提交，产生一个新的节点。


    注意1:该参数和--no-ff冲突

    注意2:该参数使用后的结果类似于在当前分支提交一个新节点。在某些情况下这个参数非常有用，例如使用Git Flow时（关于Git Flow，请参考：一个成功的Git分支模型），功能分支在进行一个功能需求的研发时，开发者可能在本地提交了大量且无意义的节点，当需要合并到develop分支时，可能仅仅需要用一个新的节点来表示这一长串节点的修改内容，这时--squash命令将会发挥作用。此外，如果功能分支的多次提交并不是琐碎而都是有意义的，使用--no-ff命令更为合适。
    --no-squash的作用正好相反。

9. -s <strategy>和 --strategy=<strategy> -s <strategy>和 --strategy=<strategy>用于指定合并的策略。默认情况如果没有指定该参数，git将按照下列情况采用默认的合并策略：

合并节点只含有单个父节点时（如采用fast-forward模式时），采用recursive策略（下文介绍）。
合并节点含有多个父节点时(如采用no-fast-forward模式时)，采用octopus策略（下文介绍）。


10. -X <option>和 --strategy-option=<option>

在-s <strategy>时指定该策略的具体参数（下文介绍）。

11. --verify-signatures, --no-verify-signatures

用于验证被合并的节点是否带有GPG签名，并在合并中忽略那些不带有GPG签名验证的节点。


12. —summary,--no-summary
和--stat与 --no-stat相似，并将在未来版本移除。

13. -q和 --quiet
静默操作，不显示合并进度信息。

14. -v和 --verbose
显示详细的合并结果信息。

15 --progress和 --no-progress
切换是否显示合并的进度信息。如果二者都没有指定，那么在标准错误发生时，将在连接的终端显示信息。请注意，并不是所有的合并策略都支持进度报告。

15-S[<keyid>]和 --gpg-sign[=<keyid>]
GPG签名。

15-m <msg>
设置用于创建合并节点时的提交信息。
如果指定了--log参数，那么commit节点的短日志将会附加在提交信息里。

15--[no-]rerere-autoupdate
rerere即reuse recorded resolution，重复使用已经记录的解决方案。它允许你让 Git 记住解决一个块冲突的方法，这样在下一次看到相同冲突时，Git 可以为你自动地解决它。

15--abort
抛弃当前合并冲突的处理过程并尝试重建合并前的状态。

### 关于合并的其他概念

1. 合并前的检测

在合并外部分支时，你应当保持自己分支的整洁，否则的话当存在合并冲突时将会带来很多麻烦。

为了避免在合并提交时记录不相关的文件，如果有任何在index所指向的HEAD节点中登记的未提交文件，git-pull和git-merge命令将会停止。

2. fast-forward合并
通常情况下分支合并都会产生一个合并节点，但是在某些特殊情况下例外。

例如调用git pull命令更新远端代码时，如果本地的分支没有任何的提交，那么没有必要产生一个合并节点。

这种情况下将不会产生一个合并节点，HEAD直接指向更新后的顶端代码，这种合并的策略就是fast-forward合并。

3. 合并细节

除了上文所提到的fast-forward合并模式以外，被合并的分支将会通过一个合并节点和当前分支绑在一起，该合并节点同时拥有合并前的当前分支顶部节点和对方分支顶部节点，共同作为父节点。
一个合并了的版本将会使所有相关分支的变化一致，包括提交节点，HEAD节点和index指针以及节点树都会被更新。只要这些节点中的文件没有重叠的地方，那么这些文件的变化都会在节点树中改动并更新保存。
如果无法明显地合并这些变化，将会发生以下的情况：

HEAD指针所指向的节点保持不变
MERGE_HEAD指针被置于其他分支的顶部
已经合并干净的路径在index文件和节点树中同时更新
对于冲突路径，index文件记录了三个版本：版本1记录了二者共同的祖先节点，版本2记录了当前分支的顶部，即HEAD，版本3记录了MERGE_HEAD。节点树中的文件包含了合并程序运行后的结果。例如三路合并算法会产生冲突。
其他方面没有任何变化。特别地，你之前进行的本地修改将继续保持原样。
如果你尝试了一个导致非常复杂冲突的合并，并想重新开始，那么可以使用git merge --abort
关于三路合并算法：
三路合并算法是用于解决冲突的一种方式，当产生冲突时，三路合并算法会获取三个节点：本地冲突的B节点，对方分支的C节点，B，C节点的共同最近祖先节点A。三路合并算法会根据这三个节点进行合并。具体过程是，B，C节点和A节点进行比较，如果B，C节点的某个文件和A节点中的相同，那么不产生冲突；如果B或C只有一个和A节点相比发生变化，那么该文件将会采用该变化了的版本；如果B和C和A相比都发生了变化，且变化不相同，那么则需要手动去合并;如果B，C都发生了变化，且变化相同，那么并不产生冲突，会自动采用该变化的版本。最终合并后会产生D节点，D节点有两个父节点，分别为B和C。

4. 合并tag

当合并一个tag时，Git总是创建一个合并的提交，即使这时能够使用fast-forward模式。该提交信息的模板预设为该tag的信息。额外地，如果该tag被签名，那么签名的检测信息将会附加在提交信息模板中。

5. 冲突是如何表示的

当产生合并冲突时，该部分会以<<<<<<<, =======和 >>>>>>>表示。在=======之前的部分是当前分支这边的情况，在=======之后的部分是对方分支的情况。

6. 如何解决冲突

在看到冲突以后，你可以选择以下两种方式：

决定不合并。这时，唯一要做的就是重置index到HEAD节点。git merge --abort用于这种情况。
解决冲突。Git会标记冲突的地方，解决完冲突的地方后使用git add加入到index中，然后使用git commit产生合并节点。
你可以用以下工具来解决冲突:
使用合并工具。git mergetool将会调用一个可视化的合并工具来处理冲突合并。
查看差异。git diff将会显示三路差异（三路合并中所采用的三路比较算法）。
查看每个分支的差异。git log --merge -p <path>将会显示HEAD版本和MERGE_HEAD版本的差异。
查看合并前的版本。git show :1:文件名显示共同祖先的版本，git show :2:文件名显示当前分支的HEAD版本，git show :3:文件名显示对方分支的MERGE_HEAD版本。


### 合并策略

Git可以通过添加-s参数来指定合并的策略。一些合并策略甚至含有自己的参数选项，通过-X<option>设置这些合并策略的参数选项。(不要忘记，合并可以在git merge和git pull命令中发生，因此该合并策略同样适用于git pull)。

1. resolve

仅仅使用三路合并算法合并两个分支的顶部节点（例如当前分支和你拉取下来的另一个分支）。这种合并策略遵循三路合并算法，由两个分支的HEAD节点以及共同子节点进行三路合并。
当然，真正会困扰我们的其实是交叉合并（criss-cross merge）这种情况。所谓的交叉合并，是指共同祖先节点有多个的情况，例如在两个分支合并时，很有可能出现共同祖先节点有两个的情况发生，这时候无法按照三路合并算法进行合并（因为共同祖先节点不唯一）。resolve策略在解决交叉合并问题时是这样处理的，这里参考《Version Control with Git》：

In criss-cross merge situations, where there is more than one possible merge basis, the resolve strategy works like this: pick one of the possible merge bases, and hope for the best. This is actually not as bad as it sounds. It often turns out that the users have been working on different parts of the code. In that case, Git detects that it's remerging some changes that are already in place and skips the duplicate changes, avoiding the conflict. Or, if these are slight changes that do cause conflict, at least the conflict should be easy for the developer to handle

这里简单翻译一下：在交叉合并的情况时有一个以上的合并基准点（共同祖先节点），resolve策略是这样工作的：选择其中一个可能的合并基准点并期望这是合并最好的结果。实际上这并没有听起来的那么糟糕。通常情况下用户修改不同部分的代码，在这种情况下，很多的合并冲突其实是多余和重复的。而使用resolve进行合并时，产生的冲突也较易于处理，真正会遗失代码的情况很少。

2. recursive
仅仅使用三路合并算法合并两个分支。和resolve不同的是，在交叉合并的情况时，这种合并方式是递归调用的，从共同祖先节点之后两个分支的不同节点开始递归调用三路合并算法进行合并，如果产生冲突，那么该文件不再继续合并，直接抛出冲突；其他未产生冲突的文件将一直执行到顶部节点。额外地，这种方式也能够检测并处理涉及修改文件名的操作。这是git合并和拉取代码的默认合并操作。
recursive合并策略有以下参数：

- ours
该参数将强迫冲突发生时，自动使用当前分支的版本。这种合并方式不会产生任何困扰情况，甚至git都不会去检查其他分支版本所包含的冲突内容这种方式会抛弃对方分支任何冲突内容。

- theirs
正好和ours相反。
theirs和ours参数都适用于合并二进制文件冲突的情况。

- patience
在这种参数下，git merge-recursive花费一些额外的时间来避免错过合并一些不重要的行（如函数的括号）。如果当前分支和对方分支的版本分支分离非常大时，建议采用这种合并方式。

- diff-algorithm=[patience|minimal|histogram|myers]
告知git merge-recursive使用不同的比较算法。

- ignore-space-change, ignore-all-space, ignore-space-at-eol
根据指定的参数来对待空格冲突。

如果对方的版本仅仅添加了空格的变化，那么冲突合并时采用我们自己的版本
如果我们的版本含有空格，但是对方的版本包含大量的变化，那么冲突合并时采用对方的版本
采用正常的处理过程

- no-renames
关闭重命名检测。

- subtree[=<path>]
该选项是subtree合并策略的高级形式，将会猜测两颗节点树在合并的过程中如何移动。不同的是，指定的路径将在合并开始时除去，以使得其他路径能够在寻找子树的时候进行匹配。（关于subtree合并策略详见下文）

- octopus
这种合并方式用于两个以上的分支，但是在遇到冲突需要手动合并时会拒绝合并。这种合并方式更适合于将多个分支捆绑在一起的情况，也是多分支合并的默认合并策略。

4.4ours
这种方式可以合并任意数量的分支，但是节点树的合并结果总是当前分支所冲突的部分。这种方式能够在替代旧版本时具有很高的效率。请注意，这种方式和recursive策略下的ours参数是不同的。

4.5subtree
subtree是修改版的recursive策略。当合并树A和树B时，如果B是A的子树，B首先调整至匹配A的树结构，而不是读取相同的节点。

4.5总结
在使用三路合并的策略时（指默认的recursive策略），如果一个文件（或一行代码）在当前分支和对方分支都产生变化，但是稍后又在其中一个分支回退，那么这种回退的变化将会在结果中体现。这一点可能会使一些人感到困惑。这是由于在合并的过程中，git仅仅关注共同祖先节点以及两个分支的HEAD节点，而不是两个分支的所有节点。因此，合并算法将会把被回退的部分认为成没有变化，这样，合并后的结果就会变为另一个分支中变化的部分。

